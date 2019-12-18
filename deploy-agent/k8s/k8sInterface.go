package k8s

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
	GetLog(d deploy, log chan string)error
}
