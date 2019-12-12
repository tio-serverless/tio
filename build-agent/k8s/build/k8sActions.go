package deploy

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"tio/build-agent/k8s/dataBus"
)

/*
Deploy Single Job In Kubernetes Cluster
*/

var kc *k8sClient

type k8sClient struct {
	client    *kubernetes.Clientset
	namespace string
}

func InitK8sClient(bus *dataBus.DataBus) (err error) {

	config, err := clientcmd.BuildConfigFromFlags("", bus.K8S.Config)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	kc = new(k8sClient)

	kc.client = clientset
	kc.namespace = bus.K8S.Namespace

	info, err := kc.client.ServerVersion()
	if err != nil {
		return err
	}

	logrus.Infof("Kubernetes Version: %s", info.String())
	return
}

// NewJob
// commenv is the common environment. Every user will use it.
func NewJob(b dataBus.BuildModel, d *dataBus.DataBus) (err error) {

	name := fmt.Sprintf("tio-%s", b.Name)

	j, err := GetJob(name)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			j = nil
		} else {
			return err
		}
	}

	if j != nil && j.Name != "" {
		if err := RemoveJob(name); err != nil {
			return err
		}
	}

	// Clear build job after 10mins.
	ttl := int32(60 * 10)
	bf := int32(1)

	job := v1.Job{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: kc.namespace,
		},
		Spec: v1.JobSpec{
			BackoffLimit:            &bf,
			TTLSecondsAfterFinished: &ttl,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: kc.namespace,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            name,
							Image:           d.BuildImage,
							ImagePullPolicy: apiv1.PullAlways,
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "dockersock",
									MountPath: "/var/run/docker.sock",
								},
							},
							Env: []apiv1.EnvVar{
								{
									Name:  "TIO_DOCKER_USER",
									Value: d.Docker.User,
								},
								{
									Name:  "TIO_DOCKER_PASSWD",
									Value: d.Docker.Passwd,
								},
							},
							Args: []string{
								"-zip", b.Address,
								"-base", d.BaseImage,
								"-control", d.Control,
								"-sid", strconv.Itoa(int(b.Sid)),
							},
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
					Volumes: []apiv1.Volume{
						{
							Name: "dockersock",
							VolumeSource: apiv1.VolumeSource{
								HostPath: &apiv1.HostPathVolumeSource{
									Path: "/var/run/docker.sock",
								},
							},
						},
					},
				},
			},
		},
	}

	j, err = kc.client.BatchV1().Jobs(kc.namespace).Create(&job)
	if err != nil {
		return err
	}

	logrus.Debugf("Job Pod Num: %d", j.Status.Succeeded)

	return nil
}

func GetJob(name string) (*v1.Job, error) {
	j, err := kc.client.BatchV1().Jobs(kc.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return j, nil
}

func RemoveJob(name string) error {
	deletePolicy := metav1.DeletePropagationForeground
	return kc.client.BatchV1().Jobs(kc.namespace).Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
