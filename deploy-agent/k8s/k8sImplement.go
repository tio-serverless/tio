package k8s

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"tio/deploy-agent/k8s/data"
)

type SimpleK8s struct {
	B           *data.B
	client      *kubernetes.Clientset
	monitorChan chan string
}

func NewSimpleK8s() *SimpleK8s {
	sk := &SimpleK8s{
		monitorChan: make(chan string, 100),
	}

	go sk.enableMonitor()

	return sk
}

func (k *SimpleK8s) enableMonitor() {
	for {
		select {
		case m := <-k.monitorChan:
			go func(m string) {
				logrus.Infof("Start Monitor %s ", m)
				stype, endpoint, err := k.deploymentIsReady(m)
				if err != nil {
					logrus.Errorf("Monitor %s Error. %s", m, err.Error())
					return
				}

				logrus.Infof("Get Deployment One Endpoint %s Type %s", endpoint, stype)
				switch stype {
				case data.GRPC:
					k.B.GetInjectGrpcChan() <- endpoint
				case data.HTTP:
					k.B.GetInjectHttpChan() <- data.HttpArch{
						Name: m,
						Url:  endpoint,
					}
				case data.TCP:
				}

			}(m)

		}
	}
}

func (k *SimpleK8s) deploymentIsReady(name string) (stype, result string, err error) {
	for i := 0; i < 10; i++ {
		d, err := k.client.AppsV1().Deployments(k.B.K.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return stype, result, err
		}

		time.Sleep(time.Duration(10*i) * time.Second)
		fmt.Printf("now %d ready %d expect %d \n", d.Status.Replicas, d.Status.ReadyReplicas, *d.Spec.Replicas)
		if d.Status.ReadyReplicas == *d.Spec.Replicas && d.Status.Replicas == *d.Spec.Replicas {
			p, err := k.client.CoreV1().Pods(k.B.K.Namespace).List(metav1.ListOptions{
				LabelSelector: fmt.Sprintf("tio-app=%s", name),
				Limit:         1,
			})
			if err != nil {
				return stype, result, err
			}

			if len(p.Items) == 0 {
				return stype, result, errors.New("Pod has zero instances")
			}
			for _, e := range p.Items[0].Spec.Containers[0].Env {
				if e.Name == "MY_SERVICE_TYPE" {
					stype = e.Value
					break
				}
			}

			switch stype {
			case data.GRPC:
				return stype, p.Items[0].Status.PodIP, nil
			case data.HTTP:
				for _, e := range p.Items[0].Spec.Containers[0].Env {
					if e.Name == "MY_SERVICE_URL" {
						result = e.Value
						break
					}
				}

				return stype, result, nil
			default:
				return stype, result, fmt.Errorf("Wrong Service Type [%s]", stype)
			}

		}
	}

	return
}

func (k *SimpleK8s) NewDeploy(d deploy) (string, error) {
	var ev []apiv1.EnvVar

	ev = append(ev, apiv1.EnvVar{
		Name: "MY_POD_IP",
		ValueFrom: &apiv1.EnvVarSource{
			FieldRef: &apiv1.ObjectFieldSelector{
				APIVersion: "v1",
				FieldPath:  "status.podIP",
			},
		},
	})
	ev = append(ev, apiv1.EnvVar{
		Name: "MY_POD_NAME",
		ValueFrom: &apiv1.EnvVarSource{
			FieldRef: &apiv1.ObjectFieldSelector{
				APIVersion: "v1",
				FieldPath:  "metadata.name",
			},
		},
	})

	for key, val := range d.Env {
		ev = append(ev, apiv1.EnvVar{
			Name:  key,
			Value: val,
		})
	}

	deployment, err := k.client.AppsV1().Deployments(k.B.K.Namespace).Create(&v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: d.Name,
			Labels: map[string]string{
				"tio-app": d.Name,
			},
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(k.B.K.Instance),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"tio-app": d.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: d.Name,
					Labels: map[string]string{
						"tio-app": d.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Volumes:        nil,
					InitContainers: nil,
					Containers: []apiv1.Container{
						{
							Name:            d.Name,
							Image:           d.Image,
							Env:             ev,
							ImagePullPolicy: apiv1.PullAlways,
							LivenessProbe: &apiv1.Probe{
								Handler: apiv1.Handler{
									Exec:    nil,
									HTTPGet: nil,
									TCPSocket: &apiv1.TCPSocketAction{
										Port: intstr.IntOrString{
											Type:   intstr.Int,
											IntVal: 80,
										},
									},
								},
								InitialDelaySeconds: 10,
								TimeoutSeconds:      5,
								PeriodSeconds:       5,
								SuccessThreshold:    1,
								FailureThreshold:    5,
							},
						},
						{
							Name:            "consul-sidecar",
							Image:           k.B.K.Sidecar,
							Env:             ev,
							ImagePullPolicy: apiv1.PullAlways,
						},
					},
					RestartPolicy:                 apiv1.RestartPolicyAlways,
					TerminationGracePeriodSeconds: int64Ptr(15),
					SecurityContext:               nil,
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	k.monitorChan <- deployment.Name

	return deployment.Name, nil
}

func (k *SimpleK8s) Scala(id string, instances int) error {
	d, err := k.client.AppsV1().Deployments(k.B.K.Namespace).Get(id, metav1.GetOptions{
	})
	if err != nil {
		return err
	}

	d.Spec.Replicas = int32Ptr(instances)

	_, err = k.client.AppsV1().Deployments(k.B.K.Namespace).Update(d)

	return err
}

func (k *SimpleK8s) Delete(id string) error {
	return nil
}

func (k *SimpleK8s) GetDeploymentInfo(name string) (v1.Deployment, error) {

	d, err := k.client.AppsV1().Deployments(k.B.K.Namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return v1.Deployment{}, err
	}

	return *d, nil
}

func (k *SimpleK8s) GetPodInfo(name string) (apiv1.Pod, error) {

	p, err := k.client.CoreV1().Pods(k.B.K.Namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("tio-app=%s", name),
		Limit:         1,
	})
	if err != nil {
		return apiv1.Pod{}, err
	}

	if len(p.Items) == 0 {
		return apiv1.Pod{}, errors.New("There are not running pod")
	}

	for _, pod := range p.Items {
		if strings.HasSuffix(pod.Name, "-sidecar") {
			return pod, nil
		}

	}
	return apiv1.Pod{}, fmt.Errorf("There are not available pod in %s", name)
}

//func (k *SimpleK8s) GetDeploymentEndpointWithName(name string) (string, error) {
//	p, err := k.client.CoreV1().Pods(k.B.K.Namespace).List(metav1.ListOptions{
//		LabelSelector: fmt.Sprintf("tio-app=%s", name),
//		Limit:         1,
//	})
//	if err != nil {
//		return "", nil
//	}
//
//	if len(p.Items) == 0 {
//		return "", errors.New("There are not running pod")
//	}
//
//	return fmt.Sprintf("%s:80", p.Items[0].Status.PodIP), nil
//}

func (k *SimpleK8s) InitClient() error {
	config, err := clientcmd.BuildConfigFromFlags("", k.B.K.Config)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	k.client = clientset

	info, err := k.client.ServerVersion()
	if err != nil {
		return err
	}

	logrus.Infof("Kubernetes Version: %s", info.String())
	return nil
}

func (k *SimpleK8s) IsHasDeploy(id string) (bool, error) {
	d, err := k.client.AppsV1().Deployments(k.B.K.Namespace).Get(id, metav1.GetOptions{
	})

	if err != nil {
		return false, err
	}

	if d.UID == "" {
		return false, nil
	}

	return true, nil
}

func (k *SimpleK8s) ReplaceDeploy(d deploy) error {
	deployClient := k.client.AppsV1().Deployments(k.B.K.Namespace)
	oldDeployment, err := deployClient.Get(d.Name, metav1.GetOptions{
	})

	if err != nil {
		return err
	}

	for i, c := range oldDeployment.Spec.Template.Spec.Containers {
		if c.Name != "consul-sidecar" {
			if d.Image != "" {
				oldDeployment.Spec.Template.Spec.Containers[i].Image = d.Image
			}
			oldDeployment.Spec.Template.Spec.Containers[i].Env = envMerge(oldDeployment.Spec.Template.Spec.Containers[i].Env, d.Env)
			continue
		}

		if c.Name == "consul-sidecar" {
			oldDeployment.Spec.Template.Spec.Containers[i].Env = envMerge(oldDeployment.Spec.Template.Spec.Containers[i].Env, d.Env)
		}
	}

	//logrus.Debugf("Update New Deployment [%v]", oldDeployment)
	_, err = deployClient.Update(oldDeployment)
	k.monitorChan <- oldDeployment.Name
	return nil
}

//  envMerge 使用env2更新env1，同时将env2中新增的key添加到env1中
func envMerge(env1 []apiv1.EnvVar, env2 map[string]string) []apiv1.EnvVar {
	if len(env2) == 0 {
		return env1
	}

	var newEnv []apiv1.EnvVar

	existKey := make(map[string]bool)
	for _, e := range env1 {
		if v, ok := env2[e.Name]; ok {
			newEnv = append(newEnv, apiv1.EnvVar{
				Name:  e.Name,
				Value: v,
			})
			existKey[e.Name] = true
		} else {
			newEnv = append(newEnv, e)
		}
	}

	for key, val := range env2 {
		if _, ok := existKey[key]; !ok {
			newEnv = append(newEnv, apiv1.EnvVar{
				Name:  key,
				Value: val,
			})
		}
	}

	return newEnv
}

func (k *SimpleK8s) GetLog(d deploy, log chan string) error {
	return k.GetDeploymentLog(d.Name, true, log)
}

func (k *SimpleK8s) Update(d deploy) error {
	return k.ReplaceDeploy(d)
}

func int32Ptr(i int) *int32 {
	i32 := int32(i)
	return &i32
}

func int64Ptr(i int) *int64 {
	_i := int64(i)
	return &_i
}
