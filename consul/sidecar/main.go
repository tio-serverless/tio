package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

var client *api.Client

func main() {

	defer func() {
		logrus.Debugf("%s deregister\n", os.Getenv("MY_POD_NAME"))
		client.Agent().ServiceDeregister(os.Getenv("MY_POD_NAME"))
		if r := recover(); r != nil {
			logrus.Errorf("panic %s", r)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		logrus.Debugf("%s deregister\n", os.Getenv("MY_POD_NAME"))
		client.Agent().ServiceDeregister(os.Getenv("MY_POD_NAME"))
		os.Exit(0)
	}()

	var err error
	for _, add := range strings.Split(os.Getenv("CONSUL_ADDRESS"), ";") {
		if err = registerMySelf(add); err != nil {
			logrus.Debugf("register error: %s", err)
			continue
		}
		break
	}

	if err != nil {
		logrus.Fatalf("register error: %s", err)
	}

	<-make(chan struct{})
}

func registerMySelf(address string) error {
	config := api.DefaultConfig()
	config.Address = address
	var err error

	client, err = api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	p, _ := strconv.Atoi(os.Getenv("MY_POD_PORT"))
	return registerService(client, p, os.Getenv("MY_POD_IP"))
}

func registerService(cli *api.Client, port int, ip string) error {

	return cli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      os.Getenv("MY_POD_NAME"),
		Name:    os.Getenv("MY_SERVICE_NAME"),
		Port:    port,
		Address: ip,
		Meta: map[string]string{
			"register": time.Now().String(),
		},
		//Weights:           nil,
		Check: &api.AgentServiceCheck{
			CheckID:                        fmt.Sprintf("%s-Check", os.Getenv("MY_POD_NAME")),
			Name:                           fmt.Sprintf("%s-Check", os.Getenv("MY_SERVICE_NAME")),
			Interval:                       "3s",
			Timeout:                        "1s",
			TCP:                            fmt.Sprintf("%s:%d", ip, port),
			DeregisterCriticalServiceAfter: "3s",
		},
	})
}
