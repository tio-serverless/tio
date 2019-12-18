package rpc

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

func GetAgentInfo(address, name string) (add string, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return add, err
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	r, err := c.GetAgentMeta(ctx, &tio_control_v1.TioAgentRequest{
		Name: name,
	})

	if err != nil {
		return add, err
	}

	if r.Code == tio_control_v1.CommonRespCode_RespFaild {
		return add, errors.New("Not found this type agent metadata ")
	}

	return r.Address, nil
}
