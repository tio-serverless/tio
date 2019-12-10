package k8s

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"tio/deploy-agent/k8s/data"
)

type SimpleK8s struct {
	B      *data.B
	client *kubernetes.Clientset
}

func (k *SimpleK8s) NewDeploy(d deploy) (string, error) {
	k.client.AppsV1().Deployments(k.B.K.Namespace).Create(&v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind: "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(k.B.K.Instance),
		},
	})
	return "", nil
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

func int32Ptr(i int) *int32 {
	i32 := int32(i)
	return &i32
}
