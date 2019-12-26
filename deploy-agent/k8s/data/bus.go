package data

import (
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type B struct {
	Log            string            `toml:"log"`
	Port           int               `toml:"port"`
	Inject         map[string]string `toml:"inject"`
	K              k8s               `toml:"k8s"`
	injectGrpcChan chan string
	injectHttpChan chan HttpArch
}

type k8s struct {
	Config    string `toml:"config"`
	Namespace string `toml:"namespace"`
	Instance  int    `toml:"instance"`
	Consul    string `toml:"consul"`
	Sidecar   string `toml:"sidecar"`
}

type MyDeploy struct {
	Meta DeployMeta `toml:"deploy"`
}

type DeployMeta struct {
	Url string `toml:"url"`
}

type HttpArch struct {
	Name string
	Url  string
}

const (
	GRPC = "0"
	HTTP = "1"
	TCP  = "2"
)

func InitBus(file string) (*B, error) {
	b := new(B)

	_, err := toml.DecodeFile(file, b)
	if err != nil {
		return nil, err
	}
	if err = isValid(b); err != nil {
		return nil, err
	}

	b.injectGrpcChan = make(chan string, 100)
	b.injectHttpChan = make(chan HttpArch, 1000)

	enableLog(b)

	output(b)
	return b, nil
}

func enableLog(b *B) {
	logrus.SetLevel(logrus.InfoLevel)
	switch b.Log {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}
}

func output(b *B) {
	logrus.Println("--------------------")
	logrus.Printf("Log: %s", b.Log)
	logrus.Printf("Port: %d", b.Port)
	logrus.Println("Inject: ")
	logrus.Printf("    GRPC: %s", b.Inject["grpc"])
	logrus.Printf("    HTTP: %s", b.Inject["http"])

	logrus.Println("K8s: ")
	logrus.Printf("    Config: %s", b.K.Config)
	logrus.Printf("    Namespace: %s", b.K.Namespace)
	logrus.Printf("    Instancers: %d", b.K.Instance)
	logrus.Printf("    Consul: %s", b.K.Consul)
	logrus.Println("--------------------")
}

func isValid(b *B) error {
	if b.K.Config == "" {
		return errors.New("K8s Config Can not Empty! ")
	}

	if b.K.Consul == "" {
		return errors.New("K8s Consul Can not Empty! ")
	}

	if b.K.Namespace == "" {
		b.K.Namespace = "default"
	}

	if b.K.Instance == 0 {
		b.K.Instance = 2
	}

	return nil
}

func (b *B) GetInjectGrpcChan() chan string {
	return b.injectGrpcChan
}

func (b *B) GetInjectHttpChan() chan HttpArch {
	return b.injectHttpChan
}
