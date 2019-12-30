package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSvcImplement_IsValidInject(t *testing.T) {
	s := &svcImplement{
		inject: map[string]string{
			"/v1/api": "svc1",
		},
	}

	actual := s.IsValidInject("/v1/api")

	assert.EqualValues(t, actual, true)
}

func Test_svcImplement_GetServiceName(t *testing.T) {
	type fields struct {
		inject map[string]string
	}
	type args struct {
		uri string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "siample",
			fields: fields{inject: map[string]string{"/v1/api": "svc1",}},
			args:   args{uri: "/v1/api"},
			want:   "svc1",
		},
		{
			name:   "noData",
			fields: fields{inject: map[string]string{"/v1/api": "svc1",}},
			args:   args{uri: "/v2/api"},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &svcImplement{
				inject: tt.fields.inject,
			}
			if got := s.GetServiceName(tt.args.uri); got != tt.want {
				t.Errorf("GetServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_svcImplement_Wait(t *testing.T) {
	type fields struct {
		srvChan chan service
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    service
		wantErr bool
	}{
		{
			name:   "withSucData",
			fields: fields{srvChan: make(chan service)},
			args: args{
				name: "svc1",
			},
			want: service{
				Name:     "svc1",
				Endpoint: "127.0.0.1:80",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &svcImplement{
				srvChan: tt.fields.srvChan,
			}
			go func(name string) {
				time.Sleep(3 * time.Second)
				s.Done(service{
					Name:     name,
					Endpoint: "127.0.0.1:80",
				})
			}(tt.args.name)
			got, err := s.Wait(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wait() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wait() got = %v, want %v", got, tt.want)
			}
		})
	}
}
