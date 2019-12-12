package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"tio/control/data"
	"tio/database/model"
)

var b *data.B
var msg chan *model.Server

func main() {
	var err error
	b, err = data.InitBus(os.Getenv("TIO_CONTROL_CONFIG"))
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	msg = make(chan *model.Server, 100)
	go restWeb()

	go deploy()
	
	startRpc()
}
