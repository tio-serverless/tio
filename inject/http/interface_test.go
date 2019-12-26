package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tio/inject/http/mock"
)

func TestInject(t *testing.T) {
	result := make(map[string][]string)

	expect := map[string][]string{
		"helloService": []string{
			"/v1/hello",
			"/v1/_ping",
		},
	}

	mockCtl := gomock.NewController(t)

	defer mockCtl.Finish()

	hi := mock.NewMockinjectHttp(mockCtl)

	hi.EXPECT().Store("helloService", []string{
		"/v1/hello",
		"/v1/_ping",
	}).Do(func(name string, urls []string) {
		result[name] = urls
	}).Return(nil)

	inject(hi, "helloService", []string{
		"/v1/hello",
		"/v1/_ping",
	})

	for key, val := range expect {
		assert.EqualValues(t, len(val), len(result[key]))
		assert.EqualValues(t, val[0], result[key][0])
		assert.EqualValues(t, val[1], result[key][1])
	}
}
