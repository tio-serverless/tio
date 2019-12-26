package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	tio_build_v1 "tio/tgrpc"
)

func sendInjectMsg(endpoint, name, injectTYpe string) error {
	conn, err := grpc.Dial(b.Inject[injectTYpe], grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect Build Service Error: %s", err.Error()))
	}

	defer conn.Close()

	c := tio_build_v1.NewInjectServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := c.NewGrpcSrv(ctx, &tio_build_v1.InjectRequest{
		Address: endpoint,
		Name:    name,
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(fmt.Sprintf("Deploy Agent Return Error. [%s]", reply.Msg))
	}

	return nil
}
