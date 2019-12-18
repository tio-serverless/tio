package rpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

func GetBuildLogs(address, name, stype string, flowing bool, logs chan string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	r, err := c.GetLogs(ctx, &tio_control_v1.TioLogRequest{
		Name:    name,
		Flowing: flowing,
		Stype:   stype,
	})

	if err != nil {
		return err
	}

	go func() {
		for {
			l, err := r.Recv()
			if err != nil {
				close(logs)
				return
			}

			logs <- l.Message
		}
	}()

	return nil
}
