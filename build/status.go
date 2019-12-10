package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"tio/tgrpc"
)

func updateStatus(address string, j *job) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect Control Service Error: %s", err.Error()))
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)

	reply, err := c.UpdateBuildStatus(context.Background(), &tio_control_v1.BuildStatus{
		User:   j.User,
		Name:   j.Name,
		Status: tio_control_v1.JobStatus(j.Status),
		Api:    j.API,
		Rate:   j.Rate,
		Image:  j.Image,
		Sid:    int32(sid),
		Raw:    raw,
		Stype:  j.SType,
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(reply.Msg)
	}

	return nil
}

func building(address string, j *job) error {
	if j.User == "" || j.Name == "" || j.Image == "" {
		return errors.New("User / Name / Image Empty! ")
	}

	j.Status = int(tio_control_v1.JobStatus_BuildIng)
	return updateStatus(address, j)
}

func succ(address string, j *job) error {
	if j.User == "" || j.Name == "" || j.Image == "" {
		return errors.New("User / Name / Image Empty! ")
	}

	j.Status = int(tio_control_v1.JobStatus_BuildSucc)
	return updateStatus(address, j)
}

func faild(address string, j *job) error {
	if j.User == "" || j.Name == "" {
		return errors.New("User / Name Empty! ")
	}

	j.Status = int(tio_control_v1.JobStatus_BuildFailed)

	return updateStatus(address, j)
}
