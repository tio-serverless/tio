package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var injectChan chan string

func main() {
	injectChan = make(chan string, 100)

	g, err := NewInject(os.Getenv("TIO_INJECT_REDIS_ADDR"), os.Getenv("TIO_INJECT_REDIS_PASSWD"))
	if err != nil {
		logrus.Fatalf("Get Grpc Inject Error %s", err.Error())
	}

	go func() {
		for {
			select {
			case j := <-injectChan:
				err = inject(g, j)
				if err != nil {
					logrus.Errorf("Inject  %s Error. %s", j, err.Error())
				}
			}
		}
	}()

	start(80)
}

func inject(i injectGrpc, add string) error {
	services, err := i.FetchServices(add)
	if err != nil {
		return err
	}

	for _, s := range services {
		m, err := i.FetchMethods(add, s)
		if err != nil {
			logrus.Errorf("Fetch Method Of %s Error. %s", s, err.Error())
			continue
		}

		methods := make([]string, len(m))

		for i := range m {
			name := strings.Split(m[i], s)
			if len(name) < 2 {
				methods[i] = ""
			} else {
				methods[i] = name[1][1:]
			}
		}

		err = i.Store(s, methods)
		if err != nil {
			logrus.Errorf("Store Method Of %s Error. %s", s, err.Error())
			continue
		}
	}

	return nil
}
