package main

import (
	"encoding/json"
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

type meta struct {
	Url       string `json:"url"`
	RouteType int    `json:"route_type"`
}

var _VERSION_, _BRANCH_ string
var client *api.Client

func enableLog() {
	switch strings.ToLower(os.Getenv("DEBUG")) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	enableLog()
	logrus.Infof("Version: %s Branch: %s", _VERSION_, _BRANCH_)

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
			logrus.Errorf("register error: %s", err)
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
	err = registerService(client, p, os.Getenv("MY_POD_IP"))
	if err != nil {
		return err
	}

	return registerKV(client)
}

func registerKV(cli *api.Client) error {
	m := meta{Url: os.Getenv("MY_SERVICE_URL")}

	//switch strings.ToLower(os.Getenv("MY_SERVICE_TYPE")) {
	//case "tcp":
	//	m.RouteType = 2
	//case "http":
	//	m.RouteType = 1
	//case "grpc":
	//	m.RouteType = 0
	//}

	m.RouteType, _ = strconv.Atoi(os.Getenv("MY_SERVICE_TYPE"))

	data, _ := json.Marshal(m)

	_, err := cli.KV().Put(&api.KVPair{
		Key:   fmt.Sprintf("tio/v1/gateway/services/%s", os.Getenv("MY_SERVICE_NAME")),
		Value: data,
	}, nil)

	return err
}

func registerService(cli *api.Client, port int, ip string) error {

	check := &api.AgentServiceCheck{
		CheckID:  fmt.Sprintf("%s-Check", os.Getenv("MY_POD_NAME")),
		Name:     fmt.Sprintf("%s-Check", os.Getenv("MY_SERVICE_NAME")),
		Interval: "3s",
		Timeout:  "1s",
		//TCP:                            fmt.Sprintf("%s:%d", ip, port),
		DeregisterCriticalServiceAfter: "3s",
	}

	switch os.Getenv("MY_SERVICE_TYPE") {
	case "2":
		check.TCP = fmt.Sprintf("%s:%d", ip, port)
	case "1":
		check.HTTP = fmt.Sprintf("http://%s:%d/_ping", ip, port)
	case "0":
		check.GRPC = fmt.Sprintf("%s:%d", ip, port)
		check.GRPCUseTLS = false
	}

	return cli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      os.Getenv("MY_POD_NAME"),
		Name:    os.Getenv("MY_SERVICE_NAME"),
		Port:    port,
		Address: ip,
		Meta: map[string]string{
			"register": time.Now().String(),
		},
		Check: check,
	})
}
