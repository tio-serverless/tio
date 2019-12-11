package k8s

import (
	"strconv"

	"github.com/BurntSushi/toml"
	"tio/database/model"
	"tio/deploy-agent/k8s/data"
)

func CreateNewDeploy(config *data.B, k MyK8s, s model.Server) (string, error) {
	d := deploy{
		Image: s.Image,
		Name:  s.Name,
	}

	var meta data.MyDeploy

	_, err := toml.Decode(s.Raw, &meta)
	if err != nil {
		return "", err
	}

	env := make(map[string]string)

	env["MY_POD_PORT"] = "80"
	env["MY_SERVICE_NAME"] = s.Name
	env["CONSUL_ADDRESS"] = config.K.Consul
	env["MY_SERVICE_TYPE"] = strconv.Itoa(s.Stype)
	env["MY_SERVICE_URL"] = meta.Meta.Url
	env["DEBUG"] = "debug"

	d.Env = env

	return k.NewDeploy(d)
}

func ScalaInstances(k MyK8s, id string, num int) error {
	return k.Scala(id, num)
}

func DeleteDeploy(k MyK8s, id string) error {
	return k.Delete(id)
}
