package rpc

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	"tio/client/model"
	tio_control_v1 "tio/tgrpc"
)

func Status(address string, uid, limit int, name string) (ss []model.Server, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return ss, err
	}

	defer conn.Close()

	c := tio_control_v1.NewControlServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	r, err := c.GetBuildStatus(ctx, &tio_control_v1.TioBuildQueryRequest{
		Uid:   int32(uid),
		Name:  name,
		Limit: int32(limit),
	})

	if err != nil {
		return ss, err
	}

	if r.Code != tio_control_v1.CommonRespCode_RespSucc {
		return ss, errors.New("No data returns")
	}

	for _, b := range r.Builds {
		ss = append(ss, model.Server{
			Name:    b.Name,
			Version: b.Version,
			Status:  tio_control_v1.JobStatus_name[int32(b.Status)],
		})
	}

	return ss, nil
}
