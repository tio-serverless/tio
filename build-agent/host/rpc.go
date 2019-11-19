package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	tio_build_v1 "tio/tgrpc"
)

func rpc() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", b.Port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	tio_build_v1.RegisterBuildServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
}

func (s *server) Build(ctx context.Context, in *tio_build_v1.Request) (*tio_build_v1.Reply, error) {

	fmt.Println(in.Address)

	err := runContainer("tio", b.Build.Image, []string{"-zip", in.Address})
	if err != nil {
		return &tio_build_v1.Reply{
			Code: -1,
			Msg:  err.Error(),
		}, nil
	}

	return &tio_build_v1.Reply{
		Code: 0,
	}, nil
}
