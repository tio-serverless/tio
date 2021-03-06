package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	deploy "tio/build-agent/k8s/build"
	"tio/build-agent/k8s/dataBus"
	tio_control_v1 "tio/tgrpc"
)

type server struct {
}

func (s server) GetLogs(in *tio_control_v1.TioLogRequest, ls tio_control_v1.LogService_GetLogsServer) error {
	logrus.Debugf("Fetch [%s] Build Logs Use Flowing  [%v] ?", in.Name, in.Flowing)
	ls.Send(&tio_control_v1.TioLogReply{
		Message: fmt.Sprintf("%s - Log", in.Name),
	})

	logs := make(chan string, 1000)

	err := deploy.GetLogs(in.Name, in.Flowing, logs)
	if err != nil {
		ls.Send(&tio_control_v1.TioLogReply{
			Message: fmt.Sprintf("Fetch Log Error. %s", err.Error()),
		})
		return err
	}

	for {
		select {
		case s, ok := <-logs:
			if !ok {
				return nil
			}

			err := ls.Send(&tio_control_v1.TioLogReply{
				Message: s,
			})

			if err != nil {
				close(logs)
				logrus.Errorf("Send log to [%s] error %s. Closed Chan", in.Name, err.Error())
				return err
			}
		}
	}
}

func (s server) Build(ctx context.Context, in *tio_control_v1.Request) (*tio_control_v1.Reply, error) {
	logrus.Debugf("New Build Request. Name: [%s] Type: [%s] Sid: [%d] Address: [%s]", in.Name, in.BuildType, in.Sid, in.Address)
	err := deploy.NewJob(dataBus.BuildModel{
		Name:      strings.ToLower(in.Name),
		Address:   in.Address,
		Sid:       in.Sid,
		BuildType: strings.ToLower(in.BuildType),
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

	ss := new(server)
	tio_control_v1.RegisterBuildServiceServer(s, ss)
	tio_control_v1.RegisterLogServiceServer(s, ss)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
