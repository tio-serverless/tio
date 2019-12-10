package data

import (
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestInitBus(t *testing.T) {
	f, err := ioutil.TempFile("", "k8s.toml")
	assert.Nil(t, err)

	_, err = f.WriteString(`log="debug"
port=80
[k8s]
	config="/config"
	consul="xxxx"
	sidecar="vikings/sidecar"`)

	assert.Nil(t, err)

	f.Close()

	b, err := InitBus(f.Name())
	assert.Nil(t, err)

	assert.EqualValues(t, "debug", b.Log)
	assert.EqualValues(t, 80, b.Port)
	assert.EqualValues(t, "/config", b.K.Config)
	assert.EqualValues(t, "xxxx", b.K.Consul)
	assert.EqualValues(t, "default", b.K.Namespace)
	assert.EqualValues(t, 2, b.K.Instance)
	assert.EqualValues(t, "vikings/sidecar", b.K.Sidecar)
}

func TestDeployMeta(t *testing.T) {
	f, err := ioutil.TempFile("", "k8s.toml")
	assert.Nil(t, err)

	_, err = f.WriteString(`log="debug"
port=80
[deploy]
	url="/v1/api"`)

	assert.Nil(t, err)

	f.Close()

	var d MyDeploy
	_, err = toml.DecodeFile(f.Name(), &d)

	assert.EqualValues(t, "/v1/api", d.Meta.Url)
}
