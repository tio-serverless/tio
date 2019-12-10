package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"tio/control/db"
	"tio/database/model"
	tio_build_v1 "tio/tgrpc"
)

func deploy() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recover From Error. [%s]", r)
			debug.PrintStack()
		}
	}()
	for {
		select {
		case s := <-msg:
			logrus.Debugf("Deploy [%s]", s.Name)
			var err error
			err = invokeDeploy(s)
			if err != nil {
				logrus.Errorf("Deploy Agent Return Error. [%s]", err.Error())
				err = db.UpdateSrvStatus(b, s.Id, model.SrvDeployFailed)
			} else {
				err = db.UpdateSrvStatus(b, s.Id, model.SrvDeploySuc)
			}

			if err != nil {
				logrus.Errorf("Invoke Deploy Agent Error. [%s]", err.Error())
			}
		}
	}
}

func invokeDeploy(s *model.Server) error {
	conn, err := grpc.Dial(b.DeployAgent, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect Build Service Error: %s", err.Error()))
	}

	defer conn.Close()

	c := tio_build_v1.NewTioDeployServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := c.NewDeploy(ctx, &tio_build_v1.DeployRequest{
		Name:   s.Name,
		Image:  s.Image,
		Config: s.Raw,
		Sid:    int32(s.Id),
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(fmt.Sprintf("Deploy Agent Return Error. [%s]", reply.Msg))
	}

	return nil
}
