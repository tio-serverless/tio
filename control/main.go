package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"tio/control/data"
)

var b *data.B

func init() {
	var err error
	b, err = data.InitBus(os.Getenv("TIO_CONTROL_CONFIG"))
	if err != nil {
		logrus.Fatalf(err.Error())
	}

}

func main() {
	restWeb()
}
