package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProxy(t *testing.T) {
	mockCtl := gomock.NewController(t)

	defer mockCtl.Finish()

	hi := NewMockdataLoader(mockCtl)

	u := "/v1/api"
	hi.EXPECT().IsValidInject(u).Return(true)
	hi.EXPECT().GetServiceName(u).Return("svc1")
	hi.EXPECT().Scala("svc1").Return(nil)
	hi.EXPECT().Wait("svc1").Return(service{
		Name:     "svc1",
		Endpoint: "127.0.0.1:80",
	}, nil)

	w := httptest.ResponseRecorder{
		Code:      200,
		HeaderMap: nil,
		Body:      nil,
		Flushed:   false,
	}

	r := &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Path: "/v1/api",
		},
		Proto: "http",
	}

	hi.EXPECT().Transfer("127.0.0.1:80", &w, r)

	Proxy(hi, &w, r)

}
