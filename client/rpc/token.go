package rpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

func Token(address, name, passwd string) (accessKey, secretKey, bucket, callBackUrl string, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	r, err := c.GetToken(ctx, &tio_control_v1.TioUserRequest{
		Name:   name,
		Passwd: passwd,
	})

	if err != nil {
		return
	}

	if r.Token == nil {
		return
	}
	return r.Token.AccessKey, r.Token.SecretKey, r.Token.Bucket, r.Token.CallBackUrl, nil
}
