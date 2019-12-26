package main

import (
	"context"

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
	cc, err := grpcurl.BlockingDial(ctx, "tcp", "172.19.64.213:80", nil, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	refClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(cc))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)

	return grpcurl.ListServices(descSource)
}

func (g *gInject) FetchMethods(s string) ([]string, error) {
	ctx := context.Background()
	cc, err := grpcurl.BlockingDial(ctx, "tcp", "172.19.64.213:80", nil, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	refClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(cc))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)
	return grpcurl.ListMethods(descSource, s)
}

func (g *gInject) Store(name string, methods []string) error {
	return nil
}

func NewInject() (*gInject, error) {
	return nil, nil
}
