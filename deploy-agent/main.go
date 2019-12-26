package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

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

	go enableInject()

	startRpc()
}

func startRpc() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", b.Port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	sk := k8s.NewSimpleK8s()
	sk.B = b
	//sk := k8s.SimpleK8s{B: b}
	if err = sk.InitClient(); err != nil {
		logrus.Fatalf("K8s Client Init Failed. [%s]", err.Error())
	}

	gs := &grcpSrv{cli: sk}

	tio_control_v1.RegisterTioDeployServiceServer(s, gs)
	tio_control_v1.RegisterLogServiceServer(s, gs)
	tio_control_v1.RegisterTioDeployCommServiceServer(s, gs)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

// enableInject 将已经部署就绪的服务Endpoint发给Inject GRPC Service
func enableInject() {
	for {
		select {
		case i := <-b.GetInjectGrpcChan():
			i = fmt.Sprintf("%s:80", i)
			if err := sendInjectMsg(i, "", "grpc"); err != nil {
				logrus.Errorf("Send GRPC Inject Error %s", err.Error())
			}
		case h := <-b.GetInjectHttpChan():
			if err := sendInjectMsg(h.Url, h.Name, "http"); err != nil {
				logrus.Errorf("Send HTTP Inject Error %s", err.Error())
			}
		}
	}
}

type grcpSrv struct {
	cli k8s.MyK8s
}

func (g grcpSrv) GetLogs(in *tio_control_v1.TioLogRequest, ls tio_control_v1.LogService_GetLogsServer) error {
	logrus.Debugf("Fetch [%s] Running Log", in.Name)
	logs := make(chan string, 1000)
	err := k8s.GetDeploymentLog(g.cli, in.Name, logs)
	if err != nil {
		logrus.Errorf("GetDeployment [%s] Log Error. %s ", in.Name, err.Error())
		return err
	}

	for {
		select {
		case l, ok := <-logs:
			if !ok {
				err := ls.Send(&tio_control_v1.TioLogReply{
					Message: "Log Output Finish!",
				})
				if err != nil {
					return err
				}
				return nil
			}

			err := ls.Send(&tio_control_v1.TioLogReply{
				Message: l,
			})
			if err != nil {
				return err
			}
		}
	}
}

func (g grcpSrv) UpdateSrvMeta(ctx context.Context, in *tio_control_v1.SrvMeta) (*tio_control_v1.TioReply, error) {
	logrus.Debugf("Update [%s] Metadata [%v]", in.Name, in.Env)
	names := strings.Split(in.Name, "-")
	if len(names) != 3 {
		return &tio_control_v1.TioReply{
			Code: tio_control_v1.CommonRespCode_RespFaild,
			Msg:  fmt.Sprintf("Name [%s] formate error. Should xxx-xxx-xxx. ", in.Name),
		}, nil
	}

	err := k8s.UpdateDeployment(g.cli, names[1], in.Env)
	if err != nil {
		return &tio_control_v1.TioReply{
			Code: tio_control_v1.CommonRespCode_RespFaild,
			Msg:  fmt.Sprintf("Update Name [%s]  error. %s ", in.Name, err.Error()),
		}, nil
	}
	return &tio_control_v1.TioReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
		Msg:  "OK",
	}, nil
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
	logrus.Debugf("New Deploy Request. Name: %s Image: %s Raw:%s", in.Name, in.Image, in.Config)
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
