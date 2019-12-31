package main

import "github.com/sirupsen/logrus"

func main() {
	mi, err := NewMonImplement()
	if err != nil {
		logrus.Fatalf("Monitor Init Error. %s", err.Error())
	}

	allSvc, err := mi.WatchProemetheus()
	if err != nil {
		logrus.Fatalf("Create Prometheus Watch Chan Error. %s", err.Error())
	}

	go func() {
		select {
		case svc := <-allSvc:
			for _, s := range svc {
				err := mi.WatchForScala(s)
				if err != nil {
					logrus.Errorf("Watch Scala %v Error. %s", s, err.Error())
				}
			}

		}
	}()

	logrus.Infof("Monitor Service Start - - - ")
	start(mi, 80)
}
