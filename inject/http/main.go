package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type arch struct {
	Name string
	Urls []string
}

var injectHttpChan chan arch

func main() {
	injectHttpChan = make(chan arch, 100)
	db, _ := strconv.Atoi(os.Getenv("TIO_INJECT_REDIS_DB"))
	g, err := NewInject(os.Getenv("TIO_INJECT_REDIS_ADDR"), os.Getenv("TIO_INJECT_REDIS_PASSWD"), db)
	if err != nil {
		logrus.Fatalf("Get Grpc Inject Error %s", err.Error())
	}

	go func() {
		for {
			select {
			case j := <-injectHttpChan:
				err = inject(g, j.Name, j.Urls)
				if err != nil {
					logrus.Errorf("Inject  %s Error. %s", j, err.Error())
				}
			}
		}
	}()

	start(80)
}

func inject(i injectHttp, name string, urls []string) error {
	return i.Store(name, urls)
}
