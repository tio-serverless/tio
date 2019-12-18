package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"tio/control/data"
	"tio/control/db"
	"tio/database/model"
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

func (s server) GetLogs(in *tio_control_v1.TioLogRequest, ls tio_control_v1.ControlService_GetLogsServer) error {
	logrus.Debugf("Fetch [%s] [%s] Logs", in.Name, in.Stype)

	logs := make(chan string, 1000)

	switch strings.ToLower(in.Stype) {
	case "build":
		if err := s.getLogFromAgent(b.BuildAgent, in.Name, in.Flowing, logs); err != nil {
			logrus.Errorf("Fetch [%s] logs error. %s", in.Name, err.Error())
			return err
		}

	case "deploy":
		if err := s.getLogFromAgent(b.DeployAgent, in.Name, in.Flowing, logs); err != nil {
			logrus.Errorf("Fetch [%s] logs error. %s", in.Name, err.Error())
			return err
		}

	default:
		return nil
	}

	for {
		select {
		case l, ok := <-logs:
			if !ok {
				ls.Send(&tio_control_v1.TioLogReply{
					Message: "Logs Chan Closed!",
				})
				return nil
			}

			ls.Send(&tio_control_v1.TioLogReply{
				Message: l,
			})
		}
	}

	return nil
}

func (s server) getLogFromAgent(address, name string, flowing bool, logs chan string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Connect Build Service Error: %s", err.Error()))
	}
	
	c := tio_control_v1.NewLogServiceClient(conn)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	reply, err := c.GetLogs(ctx, &tio_control_v1.TioLogRequest{
		Name:    name,
		Flowing: flowing,
	})

	if err != nil {
		return err
	}

	go func() {
		defer conn.Close()
		defer cancle()
		for {
			l, err := reply.Recv()
			if err != nil {
				close(logs)
				return
			}

			logrus.Debugf("Reve Log [%s]", l.Message)
			logs <- l.Message
		}
	}()

	return nil
}

func (s server) GetAgentMeta(ctx context.Context, in *tio_control_v1.TioAgentRequest) (*tio_control_v1.TioAgentReply, error) {
	switch strings.ToLower(in.Name) {
	case "build":
		return &tio_control_v1.TioAgentReply{
			Code:    tio_control_v1.CommonRespCode_RespSucc,
			Address: b.BuildAgent,
		}, nil
	case "deploy":
		return &tio_control_v1.TioAgentReply{
			Code:    tio_control_v1.CommonRespCode_RespSucc,
			Address: b.DeployAgent,
		}, nil
	default:
		return &tio_control_v1.TioAgentReply{
			Code: tio_control_v1.CommonRespCode_RespFaild,
		}, nil
	}
}

func (s server) GetBuildStatus(ctx context.Context, in *tio_control_v1.TioBuildQueryRequest) (*tio_control_v1.TioBuildQueryReply, error) {
	logrus.Infof("User [%d] Wants Query [%s] Status Limit [%d]", in.Uid, in.Name, in.Limit)
	ss, err := db.QueryUserAllSrv(s.B, int(in.Uid), int(in.Limit), in.Name)
	if err != nil {
		logrus.Errorf("Query Status Error. %s", err)
		return &tio_control_v1.TioBuildQueryReply{
			Code:   tio_control_v1.CommonRespCode_RespFaild,
			Builds: nil,
		}, nil
	}

	var bs []*tio_control_v1.BuildStatus
	for _, s := range ss {

		b := &tio_control_v1.BuildStatus{
			Name:    s.Name,
			Api:     s.Path,
			Version: s.Version,
			Status:  tio_control_v1.JobStatus(s.Status),
		}
		bs = append(bs, b)
	}
	return &tio_control_v1.TioBuildQueryReply{
		Code:   tio_control_v1.CommonRespCode_RespSucc,
		Builds: bs,
	}, nil
}

func (s server) GetToken(ctx context.Context, in *tio_control_v1.TioUserRequest) (*tio_control_v1.TioUserReply, error) {
	logrus.Debugf("User[%s] Use [%s] Wants Get Upload Token", in.Name, in.Passwd)

	_, err := db.QueryUser(s.B, in.Name, in.Passwd)
	if err != nil {
		logrus.Errorf("Query User Info Error. %s", err.Error())
		return &tio_control_v1.TioUserReply{
			Code: tio_control_v1.CommonRespCode_RespFaild,
		}, nil
	}

	return &tio_control_v1.TioUserReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
		Token: &tio_control_v1.TioToken{
			AccessKey:   s.B.Storage.AcessKey,
			SecretKey:   s.B.Storage.SecretKey,
			Bucket:      s.B.Storage.Bucket,
			CallBackUrl: s.B.Storage.CallBackUrl,
		},
	}, nil
}

func (s server) Login(ctx context.Context, in *tio_control_v1.TioUserRequest) (*tio_control_v1.TioUserReply, error) {
	logrus.Debugf("User[%s] Use [%s] Request Login", in.Name, in.Passwd)
	u, err := db.QueryUser(s.B, in.Name, in.Passwd)
	if err != nil {
		logrus.Errorf("Query User Info Error. %s", err.Error())
		return &tio_control_v1.TioUserReply{
			Code: tio_control_v1.CommonRespCode_RespFaild,
		}, nil
	}

	return &tio_control_v1.TioUserReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
		User: &tio_control_v1.TioUserInfo{
			Uid: int32(u.Id),
		},
	}, nil
}

func (s server) UpdateBuildStatus(ctx context.Context, in *tio_control_v1.BuildStatus) (*tio_control_v1.BuildReply, error) {
	logrus.Infof("user: %s name: %s image: %s rate: %d api: %s status: %d srvid: %d type: %s version: %s", in.User, in.Name, in.Image, in.Rate, in.Api, in.Status, in.Sid, in.Stype, in.Version)
	var err error

	switch in.Status {
	case tio_control_v1.JobStatus_BuildSucc:
		err = db.UpdateSrvBuildResult(b, int(in.Sid), model.SrvBuildSuc, in.Name, in.Api, in.Image, in.Raw, in.Stype, in.Version)
		if err != nil {
			logrus.Errorf("Update Srv Status Error [%s]", err)
			return &tio_control_v1.BuildReply{
				Code: -1,
				Msg:  err.Error(),
			}, nil
		}

		ns, _ := db.QuerySrvById(b, int(in.Sid))
		msg <- ns
	case tio_control_v1.JobStatus_BuildFailed:
		err = db.UpdateSrvStatus(b, int(in.Sid), model.SrvBuildFailed)
	case tio_control_v1.JobStatus_BuildIng:
		err = db.UpdateSrvStatus(b, int(in.Sid), model.SrvBuilding)
	}

	return &tio_control_v1.BuildReply{
		Code: 0,
		Msg:  "OK",
	}, nil
}
