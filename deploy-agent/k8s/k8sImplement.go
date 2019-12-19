package k8s

import (
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
	B      *data.B
	client *kubernetes.Clientset
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
		//logrus.Debugf("Find Container [%s] In [%s]", c.Name, d.Name)
		if c.Name != "coonsul-sidecar" {
			if d.Image != "" {
				oldDeployment.Spec.Template.Spec.Containers[i].Image = d.Image
			}

			//for n, e := range oldDeployment.Spec.Template.Spec.Containers[i].Env {
			//	if v, ok := d.Env[e.Name]; ok {
			//		oldDeployment.Spec.Template.Spec.Containers[i].Env[n] = apiv1.EnvVar{
			//			Name:  e.Name,
			//			Value: v,
			//		}
			//	}
			//}
			oldDeployment.Spec.Template.Spec.Containers[i].Env = envMerge(oldDeployment.Spec.Template.Spec.Containers[i].Env, d.Env)
			continue
		}

		if c.Name == "consul-sidecar" {
			//for n, e := range oldDeployment.Spec.Template.Spec.Containers[i].Env {
			//	if v, ok := d.Env[e.Name]; ok {
			//		oldDeployment.Spec.Template.Spec.Containers[i].Env[n] = apiv1.EnvVar{
			//			Name:  e.Name,
			//			Value: v,
			//		}
			//	}
			//}
			oldDeployment.Spec.Template.Spec.Containers[i].Env = envMerge(oldDeployment.Spec.Template.Spec.Containers[i].Env, d.Env)
		}
	}

	//logrus.Debugf("Update New Deployment [%v]", oldDeployment)
	_, err = deployClient.Update(oldDeployment)
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
