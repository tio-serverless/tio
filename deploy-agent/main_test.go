package main

import (
	"testing"

	"tio/deploy-agent/k8s"
)

func Test_grcpSrv_compute(t *testing.T) {
	type fields struct {
		cli k8s.MyK8s
	}
	type args struct {
		ins int
		inm float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "zero",
			fields: fields{cli: nil},
			args: args{
				ins: 1,
				inm: 0.5,
			},
			want: 1,
		},
		{
			name:   "one multiple",
			fields: fields{cli: nil},
			args: args{
				ins: 1,
				inm: 1.5,
			},
			want: 2,
		},
		{
			name:   "double",
			fields: fields{cli: nil},
			args: args{
				ins: 1,
				inm: 2.3,
			},
			want: 3,
		},
		{
			name:   "half one multiple",
			fields: fields{cli: nil},
			args: args{
				ins: 2,
				inm: 0.5,
			},
			want: 1,
		},
		{
			name:   "zero multiple",
			fields: fields{cli: nil},
			args: args{
				ins: 2,
				inm: 0,
			},
			want: 0,
		},
		{
			name:   "more than half one multiple",
			fields: fields{cli: nil},
			args: args{
				ins: 2,
				inm: 1.6,
			},
			want: 4,
		},
		{
			name:   "fix ins",
			fields: fields{cli: nil},
			args: args{
				ins: 0,
				inm: 1.6,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := grcpSrv{
				cli: tt.fields.cli,
			}
			if got := g.compute(tt.args.ins, tt.args.inm); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grcpSrv_needScala(t *testing.T) {
	type fields struct {
		cli k8s.MyK8s
	}
	type args struct {
		ins int
		inm float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int
	}{
		{
			name:   "not scala 0",
			fields: fields{cli: nil},
			args: args{
				ins: 0,
				inm: float64(0.0),
			},
			want:  false,
			want1: 0,
		},
		{
			name:   "not scala 1",
			fields: fields{cli: nil},
			args: args{
				ins: 2,
				inm: float64(1.0),
			},
			want:  false,
			want1: 0,
		},
		{
			name:   "scala 0 to 1",
			fields: fields{cli: nil},
			args: args{
				ins: 0,
				inm: float64(1.0),
			},
			want:  true,
			want1: 1,
		},
		{
			name:   "double scala",
			fields: fields{cli: nil},
			args: args{
				ins: 2,
				inm: float64(1.6),
			},
			want:  true,
			want1: 4,
		},
		{
			name:   "half one scala",
			fields: fields{cli: nil},
			args: args{
				ins: 3,
				inm: float64(0.6),
			},
			want:  true,
			want1: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := grcpSrv{
				cli: tt.fields.cli,
			}
			got, got1 := g.needScala(tt.args.ins, tt.args.inm)
			if got != tt.want {
				t.Errorf("needScala() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("needScala() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
