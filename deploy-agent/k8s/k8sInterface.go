package k8s

import (
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
)

type deploy struct {
	Name  string
	Image string
	Env   map[string]string
}

type MyK8s interface {
	NewDeploy(d deploy) (string, error)
	Scala(id string, instances int) error
	Delete(id string) error
	IsHasDeploy(id string) (bool, error)
	ReplaceDeploy(d deploy) error
	GetLog(d deploy, log chan string) error
	Update(d deploy) error
	GetDeploymentInfo(string) (v1.Deployment, error)
	GetPodInfo(string) (apiv1.Pod, error)
}
