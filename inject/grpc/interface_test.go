package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tio/inject/grpc/mock"
)

func TestInject(t *testing.T) {
	result := make(map[string][]string)

	expect := map[string][]string{
		"HelloService": []string{
			"Hello",
			"Say",
		},
	}

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	gi := mock.NewMockinjectGrpc(mockCtl)

	add := "127.0.0.1:80"
	gi.EXPECT().FetchServices(add).Return([]string{
		"HelloService",
		"EchoService",
	}, nil)
	gi.EXPECT().FetchMethods(add, "HelloService").Return([]string{
		"HelloService.Hello",
		"HelloService.Say",
	}, nil)
	gi.EXPECT().FetchMethods(add, "EchoService").Return([]string{
		"EchoService.Echo",
		"EchoService.SayEcho",
	}, nil)

	gi.EXPECT().Store("HelloService", []string{
		"Hello",
		"Say",
	}).DoAndReturn(func(name string, methods []string) {
		result[name] = methods
	}).Return(nil)

	gi.EXPECT().Store("EchoService", []string{
		"Echo",
		"SayEcho",
	}).DoAndReturn(func(name string, methods []string) {
		result[name] = methods
	}).Return(nil)

	inject(gi, "127.0.0.1:80")

	for key, val := range expect {
		assert.EqualValues(t, len(val), len(result[key]))
		assert.EqualValues(t, val[0], result[key][0])
		assert.EqualValues(t, val[1], result[key][1])
	}
}
