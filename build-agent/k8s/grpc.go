package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	deploy "tio/build-agent/k8s/build"
	"tio/build-agent/k8s/dataBus"
	tio_control_v1 "tio/tgrpc"
)

type server struct {
}

func (s server) Build(ctx context.Context, in *tio_control_v1.Request) (*tio_control_v1.Reply, error) {
	logrus.Debugf("New Build Request. Name: [%s] Sid: [%d] Address: [%s]", in.Name, in.Sid, in.Address)
	err := deploy.NewJob(dataBus.BuildModel{
		Name:    in.Name,
		Address: in.Address,
		Sid:     in.Sid,
	}, b)

	if err != nil {
		logrus.Errorf("Create Job Error. %s", err.Error())
		return &tio_control_v1.Reply{
			Code: -1,
			Msg:  err.Error(),
		}, nil
	}

	return &tio_control_v1.Reply{
		Code: 0,
		Msg:  "OK",
	}, nil
}

func start(port int) {

	p := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", p)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	tio_control_v1.RegisterBuildServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
