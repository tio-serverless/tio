package main

import (
	"os"

	"github.com/sirupsen/logrus"
	deploy "tio/build-agent/k8s/build"
	"tio/build-agent/k8s/dataBus"
)

var b *dataBus.DataBus

func main() {
	b, err := dataBus.InitBus(os.Getenv("TIO_BUILD_CONFIG"))
	if err != nil {
		logrus.Fatal(err)
	}

	err = deploy.InitK8sClient(b)
	if err != nil {
		logrus.Fatal(err)
	}

	start(b.Port)
}
