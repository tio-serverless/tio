package main

import (
	"fmt"
	"net"
	"reflect"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hs := health.NewServer()
	hs.SetServingStatus("Tio-GRPC-Service", grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(s)
	srv := &server{}

	for i := 0; i < reflect.TypeOf(srv).NumMethod(); i++ {
		method := reflect.TypeOf(srv).Method(i)
		if method.Name == "ServerInit" && method.Type.NumIn() == 1 {
			method.Func.Call([]reflect.Value{
				reflect.ValueOf(s),
			})
		}
	}

	register(s, srv)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}


