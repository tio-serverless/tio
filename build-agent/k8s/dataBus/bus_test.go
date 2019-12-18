package dataBus

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitBus(t *testing.T) {
	f, _ := ioutil.TempFile("", "build.toml")
	f.WriteString(`log="debug"
    port=80
    baseImage="ubuntu"
    control="10.0.0.1:8000"
	[build]
      grpc="tioserverless/build:grpc-v0.1.0-develop"
      http="tioserverless/build:http-v0.1.0-develop"
    [k8s]
      namespace="default"
      config="/conf/k8s"
    [docker]
      user="docker-user"
      passwd="docker-password"`)

	b, err := InitBus(f.Name())
	assert.Nil(t, err)

	assert.EqualValues(t, "debug", b.Log)
	assert.EqualValues(t, 80, b.Port)
	assert.EqualValues(t, "tioserverless/build:grpc-v0.1.0-develop", b.BuildImage["grpc"])
	assert.EqualValues(t, "tioserverless/build:http-v0.1.0-develop", b.BuildImage["http"])
	assert.EqualValues(t, "ubuntu", b.BaseImage)
	assert.EqualValues(t, "10.0.0.1:8000", b.Control)
	assert.EqualValues(t, "default", b.K8S.Namespace)
	assert.EqualValues(t, "/conf/k8s", b.K8S.Config)
	assert.EqualValues(t, "docker-user", b.Docker.User)
	assert.EqualValues(t, "docker-password", b.Docker.Passwd)
}
