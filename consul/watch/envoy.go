package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startGrpc(xds *xdsrv) {
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%s", os.Getenv("MY_GRPC_PORT")))
	envoy_api_v2.RegisterRouteDiscoveryServiceServer(grpcServer, xds)
	envoy_api_v2.RegisterClusterDiscoveryServiceServer(grpcServer, xds)

	reflection.Register(grpcServer)

	logrus.Infof("GRPC Srv Listen on: %s", fmt.Sprintf("0.0.0.0:%s", os.Getenv("MY_GRPC_PORT")))
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatal(err)
	}
}

type xdsrv struct {
	envoy  map[string][]chan *envoy_api_v2.DiscoveryResponse
	exists map[string]struct{}
}

func (x xdsrv) StreamRoutes(rs envoy_api_v2.RouteDiscoveryService_StreamRoutesServer) error {
	logrus.Debug("New Stream Routes Call")

	m, err := rs.Recv()
	if err != nil {
		logrus.Errorf("Receive Routes Error: %s", err)
		return err
	}

	id := fmt.Sprintf("%s-%s", m.Node.Id, m.Node.Cluster)

	if _, ok := x.envoy[id]; !ok {
		x.envoy[id] = []chan *envoy_api_v2.DiscoveryResponse{
			make(chan *envoy_api_v2.DiscoveryResponse),
			make(chan *envoy_api_v2.DiscoveryResponse),
		}
		x.exists[id] = struct{}{}
	}

	route := x.envoy[id][0]

	logrus.Debugf("%s Routes Chan Init !", id)

	for {
		select {
		case r := <-route:
			if err := rs.Send(r); err != nil {
				if _, ok := x.envoy[id]; ok {
					for _, c := range x.envoy[id] {
						close(c)
					}

					delete(x.envoy, id)
				}

				logrus.Errorf("Send Route Error. %s", err)
				return err
			}
		}
	}
	return nil
}

func (x xdsrv) DeltaRoutes(rs envoy_api_v2.RouteDiscoveryService_DeltaRoutesServer) error {
	return nil
}

func (x xdsrv) FetchRoutes(context.Context, *envoy_api_v2.DiscoveryRequest) (*envoy_api_v2.DiscoveryResponse, error) {
	logrus.Debugf("FetchRoutes Notice")
	return nil, nil
}

func (x xdsrv) StreamClusters(cs envoy_api_v2.ClusterDiscoveryService_StreamClustersServer) error {
	logrus.Debug("New Stream Cluster Call")

	m, err := cs.Recv()
	if err != nil {
		logrus.Errorf("Receive Cluster Error: %s", err)
		return err
	}

	id := fmt.Sprintf("%s-%s", m.Node.Id, m.Node.Cluster)
	logrus.Debugf("New Stream Cluster [%s] Call", id)
	if _, ok := x.envoy[id]; !ok {
		x.envoy[id] = []chan *envoy_api_v2.DiscoveryResponse{
			make(chan *envoy_api_v2.DiscoveryResponse),
			make(chan *envoy_api_v2.DiscoveryResponse),
		}
	}
	cluster := x.envoy[id][1]

	trigger <- struct{}{}
	logrus.Debugf("%s Cluster Need Init !", id)

	for {
		select {
		case c := <-cluster:
			if err := cs.Send(c); err != nil {
				if _, ok := x.envoy[id]; ok {
					for _, c := range x.envoy[id] {
						close(c)
					}

					delete(x.envoy, id)
				}
				logrus.Errorf("Send Cluster Error. %s", err)
				return err
			}
		}
	}
}

func (x xdsrv) DeltaClusters(cs envoy_api_v2.ClusterDiscoveryService_DeltaClustersServer) error {
	return nil

}

func (x xdsrv) FetchClusters(context.Context, *envoy_api_v2.DiscoveryRequest) (*envoy_api_v2.DiscoveryResponse, error) {
	panic("implement me")
}
