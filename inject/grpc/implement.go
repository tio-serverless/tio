package main

import (
	"context"
	"encoding/json"

	"github.com/fullstorydev/grpcurl"
	"github.com/go-redis/redis"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type gInject struct {
	redCli *redis.Client
}

func (g *gInject) FetchServices(add string) ([]string, error) {
	ctx := context.Background()
	cc, err := grpcurl.BlockingDial(ctx, "tcp", add, nil, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	refClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(cc))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)

	return grpcurl.ListServices(descSource)
}

func (g *gInject) FetchMethods(add, s string) ([]string, error) {
	ctx := context.Background()
	cc, err := grpcurl.BlockingDial(ctx, "tcp", add, nil, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	refClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(cc))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)
	return grpcurl.ListMethods(descSource, s)
}

func (g *gInject) Store(name string, methods []string) error {

	data, _ := json.Marshal(methods)
	g.redCli.Set(name, data, 0)
	return nil
}

func NewInject(add, passwd string) (*gInject, error) {

	gi := &gInject{}

	gi.redCli = redis.NewClient(&redis.Options{
		Addr:     add,
		Password: passwd,
		DB:       0,
	})

	_, err := gi.redCli.Ping().Result()
	if err != nil {
		return nil, err
	}

	return gi, nil
}
