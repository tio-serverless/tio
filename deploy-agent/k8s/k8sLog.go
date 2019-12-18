package k8s

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *SimpleK8s) getPodOfJob(name string) (podname, containername string, err error) {
	names := strings.Split(name, "-")
	if len(names) != 3 {
		return "", "", errors.New(fmt.Sprintf("Deployment Name Wrong[%s]. Should xxx-xx-xxx. ", name))
	}

	containername = names[1]
	selector := fmt.Sprintf("tio-app=%s", containername)

	logrus.Debugf("Select Pod via %s", selector)

	p, err := k.client.CoreV1().Pods(k.B.K.Namespace).List(metav1.ListOptions{
		LabelSelector: selector,
	})

	if err != nil {
		return
	}

	l := len(p.Items)
	if l == 0 {
		err = errors.New("Can not find the pod of this job. ")
		return
	}

	return p.Items[l-1].Name, containername, nil
}

// GetLogs
// 读取指定Job日志，并将日志写入到chan中
func (k *SimpleK8s) GetDeploymentLog(name string, flowing bool, logs chan string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Fetch %s logs error. %s Has recoverd.", name, r)
		}
	}()

	pod, container, err := k.getPodOfJob(name)
	if err != nil {
		return
	}

	logrus.Infof("Find pod %s of job %s", pod, name)

	line := int64(1000)
	req := k.client.CoreV1().Pods(k.B.K.Namespace).GetLogs(pod, &apiv1.PodLogOptions{
		Container: container,
		TailLines: &line,
		Follow:    flowing,
	})

	podLogs, err := req.Stream()
	if err != nil {
		return errors.New("error in opening stream")
	}

	go func() {
		defer func() {
			podLogs.Close()
			if r := recover(); r != nil {
				logrus.Errorf("panic %s recover", r)
			}
		}()

		data := make([]byte, 1024)

		for {
			n, err := podLogs.Read(data)
			if err != err {
				logrus.Errorf("read %s log error. %s", name, err.Error())
				logs <- string(data[:n])
				return
			}
			if n == 0 {
				return
			}
			logs <- string(data[:n])
		}
	}()

	return
}
