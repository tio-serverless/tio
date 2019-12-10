package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"tio/database/model"
	"tio/deploy-agent/k8s"
	"tio/deploy-agent/k8s/data"
	tio_control_v1 "tio/tgrpc"
)

var b *data.B

func main() {

	var err error
	b, err = data.InitBus(os.Getenv("TIO_DEPLOY_CONFIG"))
	if err != nil {
		logrus.Fatalf("Bus Init Error. %s", err.Error())
	}

	startRpc()
}

func startRpc() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", b.Port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	sk := k8s.SimpleK8s{B: b}
	if err = sk.InitClient(); err != nil {
		logrus.Fatalf("K8s Client Init Failed. [%s]", err.Error())
	}

	tio_control_v1.RegisterTioDeployServiceServer(s, &grcpSrv{
		cli: &sk,
	})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

type grcpSrv struct {
	cli k8s.MyK8s
}

func (g grcpSrv) ScalaDeploy(ctx context.Context, in *tio_control_v1.DeployRequest) (*tio_control_v1.TioReply, error) {
	err := k8s.ScalaInstances(g.cli, in.Name, int(in.InstanceNum))
	if err != nil {
		return &tio_control_v1.TioReply{
			Code: -1,
			Msg:  err.Error(),
		}, nil
	}

	return &tio_control_v1.TioReply{
		Code: 0,
		Msg:  "OK",
	}, nil
}

func (g grcpSrv) NewDeploy(ctx context.Context, in *tio_control_v1.DeployRequest) (*tio_control_v1.TioReply, error) {
	id, err := k8s.CreateNewDeploy(b, g.cli, model.Server{
		Id:    int(in.Sid),
		Name:  in.Name,
		Image: in.Image,
		Raw:   in.Config,
		Stype: int(in.Stype),
	})

	if err != nil {
		return &tio_control_v1.TioReply{
			Code: -1,
			Msg:  err.Error(),
		}, nil
	}

	return &tio_control_v1.TioReply{
		Code: 0,
		Msg:  id,
	}, nil
}

func (g grcpSrv) UpdateSrvMeta(context.Context, *tio_control_v1.SrvMeta) (*tio_control_v1.TioReply, error) {
	panic("implement me")
}