package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

type svcImplement struct {
	ri *redis.Client
	//inject key=uri, value=service name
	inject  map[string]string
	srvChan chan service
}

func (s *svcImplement) InjectSync(context.Context, *tio_control_v1.ProxyEndpointRequest) (*tio_control_v1.TioReply, error) {
	err := s.LoadInjectData()
	if err != nil {
		logrus.Errorf("Sync Error. %s", err.Error())
	}

	return &tio_control_v1.TioReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
		Msg:  "OK",
	}, nil
}

func (s *svcImplement) NewEndpoint(ctx context.Context, in *tio_control_v1.ProxyEndpointRequest) (*tio_control_v1.TioReply, error) {
	logrus.Infof("A Service Has Create. %s Endpoint: %s", in.Name, in.Endpoint)
	s.Done(service{
		Name:     in.Name,
		Endpoint: in.Endpoint,
	})

	return &tio_control_v1.TioReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
		Msg:  "OK",
	}, nil
}

func (s *svcImplement) IsValidInject(uri string) bool {
	_, exist := s.inject[uri]
	return exist
}

func (s *svcImplement) GetServiceName(uri string) string {
	return s.inject[uri]
}

func newSI() (*svcImplement, error) {
	si := new(svcImplement)

	err := si.redis()
	if err != nil {
		return nil, err
	}

	err = si.LoadInjectData()
	if err != nil {
		return nil, err
	}

	si.srvChan = make(chan service, 100)

	return si, nil
}

func (s *svcImplement) redis() error {
	db, _ := strconv.Atoi(os.Getenv("TIO_PROXY_REDIS_DB"))
	add := os.Getenv("TIO_PROXY_REDIS_ADDR")
	passwd := os.Getenv("TIO_PROXY_REDIS_PASSWD")
	logrus.Infof("Connect Redis %s:%d use %s", add, db, passwd)

	s.ri = redis.NewClient(&redis.Options{
		Addr:     add,
		Password: passwd,
		DB:       db,
	})

	_, err := s.ri.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func (s *svcImplement) LoadInjectData() error {
	if s.ri == nil {
		if err := s.redis(); err != nil {
			logrus.Fatalf("Connect Redis Error. %s", err.Error())
		}

		s.inject = make(map[string]string)
	}

	iter := s.ri.Scan(0, "", 0).Iterator()
	for iter.Next() {
		service := iter.Val()
		r, _ := s.ri.Get(service).Result()

		var uri []string

		if err := json.Unmarshal([]byte(r), &uri); err != nil {
			logrus.Errorf("Unmarshal [%s] Error. %s", r, err.Error())
			continue
		}

		for _, u := range uri {
			s.inject[u] = service
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}

func (s *svcImplement) Scala(name string) error {
	conn, err := grpc.Dial(os.Getenv("TIO_MONITOR_ADDR"), grpc.WithInsecure())
	//conn, err := grpc.Dial("172.31.0.87:80", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()

	c := tio_control_v1.NewMonitorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Scala(ctx, &tio_control_v1.MonitorScalaRequest{
		Name: name,
		Num:  2,
	})

	if err != nil {
		return err
	}

	if r.Code != tio_control_v1.CommonRespCode_RespSucc {
		return fmt.Errorf("Monitor Scala Error %s.", r.Msg)
	}

	return nil
}

func (s *svcImplement) Wait(name string) (service, error) {
	srv := <-s.srvChan

	return srv, nil
}

func (s *svcImplement) Done(srv service) error {
	s.srvChan <- srv
	return nil
}
