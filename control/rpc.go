package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"tio/control/data"
	tio_control_v1 "tio/tgrpc"
)

func startRpc() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", b.RpcProt))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	tio_control_v1.RegisterControlServiceServer(s, &server{B: b})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	B *data.B
}

func (s server) UpdateBuildStatus(ctx context.Context, in *tio_control_v1.BuildStatus) (*tio_control_v1.BuildReply, error) {
	logrus.Infof("user: %s name: %s image: %s rate: %d api: %s status: %d", in.User, in.Name, in.Status)

	return &tio_control_v1.BuildReply{
		Code: 0,
		Msg:  "OK",
	}, nil
}

func (s server) GetBuildStatus(context.Context, *tio_control_v1.BuildStatus) (*tio_control_v1.BuildReply, error) {
	panic("implement me")
}
