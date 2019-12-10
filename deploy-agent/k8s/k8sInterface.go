package k8s

type deploy struct {
	Image string
	Env   map[string]string
}

type MyK8s interface {
	NewDeploy(d deploy) (string, error)
	Scala(id string, instances int) error
	Delete(id string) error
}



