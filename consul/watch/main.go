package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/sirupsen/logrus"
)

var _VERSION_, _BRANCH_ string
var cs chan map[string][]service
var cluster, routes map[string]chan *envoy_api_v2.DiscoveryResponse

var envoyMeta map[string][]chan *envoy_api_v2.DiscoveryResponse

var trigger chan struct{}

func main() {
	enableLog()
	cs = make(chan map[string][]service, 100)
	trigger = make(chan struct{})

	routes = make(map[string]chan *envoy_api_v2.DiscoveryResponse)
	cluster = make(map[string]chan *envoy_api_v2.DiscoveryResponse)

	envoyMeta = make(map[string][]chan *envoy_api_v2.DiscoveryResponse)

	logrus.Infof("Version: %s Branch: %s", _VERSION_, _BRANCH_)

	go watchConsul()

	srvEnvoy()
}

func srvEnvoy() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("srvEnvoy Painc Occurd. %s", r)
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	xds := &xdsrv{
		envoy:  envoyMeta,
		exists: make(map[string]struct{}),
	}

	startGrpc(xds)
}

func watchConsul() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("watchConsul Painc Occurd. %s", r)
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	cli, err := initClient()
	if err != nil {
		logrus.Fatalf("Consul Client Init Error: %s", err.Error())
	}

	err = cli.routeInit()
	if err != nil {
		logrus.Fatalf("Route Init Error. %s  Quit!", err.Error())
	}

	cli.clusterInit()

	for name, detail := range cli.routes {
		logrus.Debugf("service [%s] -> [%s]: ", name, detail[0].url)
		for _, d := range detail {
			logrus.Debugf("    %s", d.endpoint)
		}
	}

	for route, cluster := range cli.defaultCluster {
		t := ""
		switch route {
		case HTTPRoute:
			t = "Http"
		case GRPCRoute:
			t = "Grpc"
		default:
			t = "Tcp"
		}

		logrus.Debugf("%s service route cluster: %s", t, cluster)
	}

	go func(cli *client) {
		for {
			select {
			case <-trigger:
				send2Envoy(cli)
			}
		}
	}(cli)

	go cli.watchChecks(cli.cc)
	go cli.watchKVEvents(cli.cc)

	watch(cli, cli.cc)

}

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

func send2Envoy(cli *client) {
	allClusters := cli.routeConvert2ClusterDiscoveryResponse(cli.routes)
	allRoutes := cli.routeConvert2RouterDiscoveryResponse(cli.routes)

	for id, cs := range envoyMeta {
		go func(id string, cs []chan *envoy_api_v2.DiscoveryResponse) {
			cs[1] <- allClusters
			cs[0] <- allRoutes
			logrus.Debugf("Envoy [%s] Configure Reflash", id)
		}(id, cs)
	}

	logrus.Debugf("Envoy Cluster [%d] Reflash Complete", len(envoyMeta))
}
