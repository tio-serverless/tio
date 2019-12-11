package dataBus

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type BuildModel struct {
	Name    string
	Address string
	Sid     int32
}

type docker struct {
	User   string
	Passwd string
}

type DataBus struct {
	Port       int     `toml:"port"`
	K8S        k8sConf `toml:"k8s"`
	Docker     docker  `toml:"docker"`
	Log        string  `toml:"log"`
	BuildImage string  `toml:"buildImage"` //构建服务的基础镜像
	BaseImage  string  `toml:"baseImage"`  //运行服务的基础镜像
	Control    string  `toml:"control"`
}

/*
k8sConf

Kubernetes metadata.
*/
type k8sConf struct {
	// Namespace
	Namespace string `toml:"namespace"`
	// Config file path, if use token, this property can empty
	Config string `toml:"config"`
}

func InitBus(file string) (b *DataBus, err error) {

	if file == "" {
		file = "k8s.toml"
	}

	b = new(DataBus)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	_, err = toml.Decode(string(data), b)
	if err != nil {
		return
	}

	//b.LanguageRuntime = drawOffImg(b.Language)

	if err = isValid(b); err != nil {
		return
	}

	enableLog(b)
	debug(b)

	return
}

// drawOffImg
// Convert Language string to struct.
func drawOffImg(lan map[string][]string) map[string]map[string]string {
	runtime := make(map[string]map[string]string)

	for key, value := range lan {
		image := make(map[string]string)

		for _, v := range value {
			if strings.Contains(v, ":") {
				_v := strings.Split(v, ":")
				image[_v[1]] = _v[0]
			} else {
				image["latest"] = v
			}
		}

		runtime[key] = image
	}

	return runtime
}

func isValid(bus *DataBus) error {

	if bus.K8S.Config == "" {
		return errors.New("No Valid Kubernetes config file! ")
	}

	if bus.K8S.Namespace == "" {
		bus.K8S.Namespace = "default"
	}

	if bus.Port == 0 {
		bus.Port = 80
	}

	return nil
}

func enableLog(b *DataBus) {
	switch b.Log {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func debug(bus *DataBus) {
	logrus.Debug("DATA-BUS")
	logrus.Debug("*************************************")
	logrus.Debugf("Listen on: %d", bus.Port)
	logrus.Debugf("Control: %s", bus.Control)
	logrus.Debug("Kubernetes: ")
	logrus.Debugf("  Namespace: %s", bus.K8S.Namespace)
	logrus.Debugf("  Config: %s", bus.K8S.Config)
	logrus.Debug("Docker: ")
	logrus.Printf("  User: %s*****", bus.Docker.User[:2])
	logrus.Printf("  Passwd: %s*****", bus.Docker.Passwd[:2])
	logrus.Debug("*************************************")
}
