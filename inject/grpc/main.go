package main

import (
	"github.com/sirupsen/logrus"
)

var injectChan chan string

func main() {
	injectChan = make(chan string, 100)

	g, err := NewInject()
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

		err = i.Store(s, m)
		if err != nil {
			logrus.Errorf("Store Method Of %s Error. %s", s, err.Error())
			continue
		}
	}

	return nil
}
