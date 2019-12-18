package deploy

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getPodOfJob(jobname string) (podname string, err error) {
	selector := fmt.Sprintf("job-name=tio-%s", jobname)

	logrus.Debugf("Select Pod via %s", selector)

	p, err := kc.client.CoreV1().Pods(kc.namespace).List(metav1.ListOptions{
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

	return p.Items[l-1].Name, nil
}

// GetLogs
// 读取指定Job日志，并将日志写入到chan中
func GetLogs(jobname string, flowing bool, logs chan string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Fetch %s logs error. %s Has recoverd.", jobname, r)
		}
	}()

	pod, err := getPodOfJob(jobname)
	if err != nil {
		return
	}

	logrus.Infof("Find pod %s of job %s", pod, jobname)

	line := int64(1000)
	req := kc.client.CoreV1().Pods(kc.namespace).GetLogs(pod, &apiv1.PodLogOptions{
		TailLines: &line,
		Follow:    flowing,
	})

	podLogs, err := req.Stream()
	if err != nil {
		return errors.New("error in opening stream")
	}

	defer podLogs.Close()

	for {
		data := make([]byte, 1024)
		n, err := podLogs.Read(data)
		if err != nil {
			logrus.Errorf("read %s log error. %s", jobname, err.Error())
			logs <- string(data[:n])
			break
		}
		logrus.Debugf("Read Log %s", string(data[:n]))
		logs <- string(data[:n])
	}

	return
}
