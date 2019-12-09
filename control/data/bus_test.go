package data

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitBus(t *testing.T) {
	f, err := ioutil.TempFile("", "tio.toml")
	assert.Nil(t, err)

	_, err = f.WriteString(`log="debug"
rest_port=80
rpc_port=8000
build_agent_address="build.agent.tio:80"`)

	assert.Nil(t, err)

	f.Close()

	b, err := InitBus(f.Name())
	assert.Nil(t, err)

	assert.EqualValues(t, b.Log, "debug")
	assert.EqualValues(t, b.RestPort, 80)
	assert.EqualValues(t, b.RpcProt, 8000)
	assert.EqualValues(t, b.BuildAgent, "build.agent.tio:80")

	os.Setenv("TIO_CONTROL_S_AKEY", "akey")
	os.Setenv("TIO_CONTROL_S_SKEY", "skey")
	os.Setenv("TIO_CONTROL_S_DOMAIN", "domain")

	b, err = InitBus(f.Name())
	assert.Nil(t, err)

	assert.EqualValues(t, b.Storage.AcessKey, "akey")
	assert.EqualValues(t, b.Storage.SecretKey, "skey")
	assert.EqualValues(t, b.Storage.Domain, "domain")
}
