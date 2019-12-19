package rpc

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

func Update(address, name string, env map[string]string) (err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	r, err := c.UpdateServerMetadata(ctx, &tio_control_v1.SrvMeta{
		Name: name,
		Env:  env,
	})

	if err != nil {
		return err
	}

	if r.Code != tio_control_v1.CommonRespCode_RespSucc {
		return errors.New(r.Msg)
	}

	return nil
}
