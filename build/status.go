package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

func updateStatus(address, user, name string, status tio_control_v1.JobStatus) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect Control Service Error: %s", err.Error()))
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)

	reply, err := c.UpdateBuildStatus(context.Background(), &tio_control_v1.BuildStatus{
		User:   user,
		Name:   name,
		Status: status,
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(reply.Msg)
	}

	return nil
}

func succ(adress, user, name string) error {
	if user == "" || name == "" {
		return errors.New("User / Name Empty! ")
	}
	return updateStatus(adress, user, name, tio_control_v1.JobStatus_BuildSucc)
}

func faild(adress, user, name string) error {
	if user == "" || name == "" {
		return errors.New("User / Name Empty! ")
	}

	return updateStatus(adress, user, name, tio_control_v1.JobStatus_BuildFailed)
}
