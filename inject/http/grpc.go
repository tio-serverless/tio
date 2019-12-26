package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	tio_control_v1 "tio/tgrpc"
)

type server struct{}

func (s server) NewGrpcSrv(ctx context.Context, in *tio_control_v1.InjectRequest) (*tio_control_v1.TioReply, error) {
	injectHttpChan <- arch{
		Name: in.Name,
		Urls: []string{
			in.Address,
		},
	}

	return &tio_control_v1.TioReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
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

	ss := new(server)
	tio_control_v1.RegisterInjectServiceServer(s, ss)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
