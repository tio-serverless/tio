package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	tio_control_v1 "tio/tgrpc"
)

func Test_scala(t *testing.T) {

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mig := NewMockmonitorInterface(mockCtl)

	mig.EXPECT().Sacla("svc1", 2.0).Return(nil)
	mig.EXPECT().WaitScala("svc1").Return("127.0.0.1:80",nil)

	type args struct {
		mi  monitorInterface
		ctx context.Context
		in  *tio_control_v1.MonitorScalaRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *tio_control_v1.TioReply
		wantErr bool
	}{
		{
			name: "svc1",
			args: args{
				mi:  mig,
				ctx: context.Background(),
				in: &tio_control_v1.MonitorScalaRequest{
					Name: "svc1",
					Num:  2.0,
				},
			},
			want: &tio_control_v1.TioReply{
				Code: 0,
				Msg:  "127.0.0.1:80",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scala(tt.args.mi, tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("scala() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scala() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ploy(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mig := NewMockmonitorInterface(mockCtl)

	mig.EXPECT().UpdatePloy("svc1", 10)

	type args struct {
		mi  monitorInterface
		ctx context.Context
		in  *tio_control_v1.MonitorScalaRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *tio_control_v1.TioReply
		wantErr bool
	}{
		{
			name: "svc1",
			args: args{
				mi:  mig,
				ctx: context.Background(),
				in: &tio_control_v1.MonitorScalaRequest{
					Name: "svc1",
					Num:  10,
				},
			},
			want: &tio_control_v1.TioReply{
				Code: tio_control_v1.CommonRespCode_RespSucc,
				Msg:  "OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ploy(tt.args.mi, tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ploy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ploy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monImplement_serviceSala(t *testing.T) {

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	pig := NewMockprometheusInterface(mockCtl)

	pig.EXPECT().QueryAllCluster().Return([]string{
		"svc1_cluster",
		"svc3_cluster",
	}, nil)

	pig.EXPECT().QueryRange("svc1_cluster", StepMinute, 2).Return(100, nil)
	//pig.EXPECT().QueryRange("svc3_cluster", StepMinute, 2).Return(105, nil)

	type fields struct {
		proxyService   []string
		envoyService   []string
		controlService string
		deployService  string
		ploy           map[string]int
		proImp         prometheusInterface
		wait           map[string]chan struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []envoyTraffic
	}{
		{
			name: "twoSvc",
			fields: fields{
				ploy:   map[string]int{"svc1": 20, "svc2": 40},
				proImp: pig,
			},
			want: []envoyTraffic{
				{
					Name:         "svc1",
					TrafficCount: 100,
				},
				//{
				//	Name:         "svc2",
				//	TrafficCount: 105,
				//},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := monImplement{
				proxyService: tt.fields.proxyService,
				controlService: tt.fields.controlService,
				deployService:  tt.fields.deployService,
				ploy:           tt.fields.ploy,
				proImp:         tt.fields.proImp,
			}

			if got := m.serviceSala(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceSala() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monImplement_NeedScala(t *testing.T) {
	type fields struct {
		proxyService   []string
		envoyService   []string
		controlService string
		deployService  string
		ploy           map[string]int
		proImp         prometheusInterface
		wait           map[string]chan struct{}
	}
	type args struct {
		traffic envoyTraffic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  float64
	}{
		{
			name: "not need scala",
			fields: fields{
				ploy: map[string]int{"svc1": 30},
			},
			args: args{traffic: envoyTraffic{
				Name:         "svc1",
				TrafficCount: 30,
			}},
			want:  false,
			want1: 0,
		},
		{
			name: "2X need scala",
			fields: fields{
				ploy: map[string]int{"svc1": 30},
			},
			args: args{traffic: envoyTraffic{
				Name:         "svc1",
				TrafficCount: 75,
			}},
			want:  true,
			want1: 2,
		},
		{
			name: "3X need scala",
			fields: fields{
				ploy: map[string]int{"svc1": 30},
			},
			args: args{traffic: envoyTraffic{
				Name:         "svc1",
				TrafficCount: 100,
			}},
			want:  true,
			want1: 3,
		},
		{
			name: "scala to 0",
			fields: fields{
				ploy: map[string]int{"svc1": 30},
			},
			args: args{traffic: envoyTraffic{
				Name:         "svc1",
				TrafficCount: 0,
			}},
			want:  true,
			want1: 0,
		},
		{
			name: "scala to 0.5",
			fields: fields{
				ploy: map[string]int{"svc1": 30},
			},
			args: args{traffic: envoyTraffic{
				Name:         "svc1",
				TrafficCount: 10,
			}},
			want:  true,
			want1: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := monImplement{
				proxyService:   tt.fields.proxyService,
				controlService: tt.fields.controlService,
				deployService:  tt.fields.deployService,
				ploy:           tt.fields.ploy,
				proImp:         tt.fields.proImp,
			}
			got, got1 := m.NeedScala(tt.args.traffic)
			if got != tt.want {
				t.Errorf("NeedScala() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NeedScala() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
