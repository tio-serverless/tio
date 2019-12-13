package rpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

func Login(address, name, passwd string) (uid int, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return uid, err
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	r, err := c.Login(ctx, &tio_control_v1.TioUserRequest{
		Name:   name,
		Passwd: passwd,
	})

	if err != nil {
		return uid, err
	}

	return int(r.User.Uid), nil
}
