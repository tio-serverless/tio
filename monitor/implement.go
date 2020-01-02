package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	tio_control_v1 "tio/tgrpc"
)

type monImplement struct {
	proxyService      []string
	prometheusService string
	controlService    string
	deployService     string
	ploy              map[string]int
	proImp            prometheusInterface
	//wait           map[string]chan struct{}
}

func NewMonImplement() (*monImplement, error) {
	mi := new(monImplement)

	if os.Getenv("TIO_MONITOR_PROMETHEUS_ADDR") == "" {
		return nil, fmt.Errorf("TIO_MONITOR_PROMETHEUS_ADDR Empty!")
	}

	if os.Getenv("TIO_MONITOR_CONTROL_ADDR") == "" {
		return nil, fmt.Errorf("TIO_MONITOR_CONTROL_ADDR Empty!")
	}

	if os.Getenv("TIO_MONITOR_PROXY_ADDR") == "" {
		return nil, fmt.Errorf("TIO_MONITOR_PROXY_ADDR Empty!")
	}

	if os.Getenv("TIO_MONITOR_DEPLOY_ADDR") == "" {
		return nil, fmt.Errorf("TIO_MONITOR_DEPLOY_ADDR Empty!")
	}

	mi.prometheusService = os.Getenv("TIO_MONITOR_PROMETHEUS_ADDR")
	mi.proxyService = strings.Split(os.Getenv("TIO_MONITOR_PROXY_ADDR"), ";")
	mi.controlService = os.Getenv("TIO_MONITOR_CONTROL_ADDR")
	mi.deployService = os.Getenv("TIO_MONITOR_DEPLOY_ADDR")

	mi.ploy = make(map[string]int)
	client, err := api.NewClient(api.Config{
		Address: mi.prometheusService,
	})

	if err != nil {
		return nil, err
	}
	mi.proImp = prometheusImplement{cli: client}
	err = mi.InitPloy()
	if err != nil {
		return nil, fmt.Errorf("Ploy Init Error %s", err.Error())
	}

	return mi, nil
}

func (m monImplement) Scala(ctx context.Context, in *tio_control_v1.MonitorScalaRequest) (*tio_control_v1.TioReply, error) {
	return scala(m, ctx, in)
}

func (m monImplement) Ploy(ctx context.Context, in *tio_control_v1.MonitorScalaRequest) (*tio_control_v1.TioReply, error) {
	return ploy(m, ctx, in)
}

func scala(mi monitorInterface, ctx context.Context, in *tio_control_v1.MonitorScalaRequest) (*tio_control_v1.TioReply, error) {
	err := mi.Sacla(in.Name, float64(in.Num))
	endpoint, err := mi.WaitScala(in.Name)

	if err != nil {
		return &tio_control_v1.TioReply{
			Code: tio_control_v1.CommonRespCode_RespFaild,
			Msg:  err.Error(),
		}, nil
	}
	return &tio_control_v1.TioReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
		Msg:  endpoint,
	}, nil
}

func ploy(mi monitorInterface, ctx context.Context, in *tio_control_v1.MonitorScalaRequest) (*tio_control_v1.TioReply, error) {
	mi.UpdatePloy(in.Name, int(in.Num))

	return &tio_control_v1.TioReply{
		Code: tio_control_v1.CommonRespCode_RespSucc,
		Msg:  "OK",
	}, nil
}

func (m monImplement) WatchProemetheus() (chan []envoyTraffic, error) {
	traffic := make(chan []envoyTraffic, 100)

	go m.watchPrometheus(traffic)

	return traffic, nil
}

func (m monImplement) watchPrometheus(traffic chan []envoyTraffic) {
	logrus.Infof("Prometheus Watch Start - - -")
	c := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-c.C:
			etfs := m.serviceSala()
			if len(etfs) > 0 {
				traffic <- etfs
			}
		}
	}
}

// serviceSala 判断当前是否有需要scala的服务
func (m monImplement) serviceSala() []envoyTraffic {
	var etfs []envoyTraffic
	clusters, err := m.proImp.QueryAllCluster()
	if err != nil {
		logrus.Errorf("Query Cluster Error %s", err.Error())
		return etfs
	}

	ploy := m.GetPloy()

	for _, c := range clusters {
		if ploy[c] > 0 {
			//	如果存在策略，后续判断才有意义
			connectCount, err := m.queryConnectInMinuteRange(c, 1)
			if err != nil {
				logrus.Errorf("Query Cluster %s Connect Error %s", c, err.Error())
				continue
			}

			etfs = append(etfs, envoyTraffic{
				Name:         c,
				TrafficCount: connectCount,
			})

		}
	}

	return etfs
}

func (m monImplement) WatchForScala(traffic envoyTraffic) error {
	isNeedScala, instances := m.NeedScala(traffic)

	if isNeedScala {
		err := m.Sacla(traffic.Name, instances)
		if err != nil {
			return fmt.Errorf(" Cluster %s Scala Error %s", traffic.Name, err.Error())
		}

		go func(name string) {
			_, err := m.WaitScala(name)
			if err != nil {
				logrus.Errorf("Wait Cluster %s Scala Error %s", traffic.Name, err.Error())
			}

		}(traffic.Name)

	}

	return nil
}

func (m monImplement) Sacla(name string, num float64) error {
	conn, err := grpc.Dial(m.deployService, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return fmt.Errorf("Dial DeployAgent Error. %s", err.Error())
	}

	c := tio_control_v1.NewTioDeployServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	reply, err := c.ScalaDeploy(ctx, &tio_control_v1.DeployRequest{
		Name:        name,
		InstanceNum: int32(num),
	})

	if err != nil {
		return err
	}

	if reply.Code != tio_control_v1.CommonRespCode_RespSucc {
		return fmt.Errorf("Scala %s %f Error. %s", name, num, err.Error())
	}

	//m.wait[name] = make(chan struct{})

	return nil
}

func (m monImplement) WaitScala(name string) (string, error) {
	conn, err := grpc.Dial(m.deployService, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return "", fmt.Errorf("Dial DeployAgent Error. %s", err.Error())
	}

	c := tio_control_v1.NewTioDeployServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	i := 0
	for {
		logrus.Infof("Query %s Scala Stauts %d times", name, i)

		reply, err := c.DeployInfo(ctx, &tio_control_v1.DeployRequest{
			Name: name,
		})

		if err != nil {
			return "", err
		}

		if reply.Code == tio_control_v1.CommonRespCode_RespSucc {
			return reply.Msg, nil
		}

		i++
		time.Sleep(5 * time.Second)
	}

}

//func (m monImplement) IsScalaSucc(name string) (bool, error) {
//	//<-m.wait[name]
//	return true, nil
//}

func (m monImplement) InvokeDeployService(name string, num float64) error {
	return m.Sacla(name, num)
}

func (m monImplement) InitPloy() error {
	conn, err := grpc.Dial(m.controlService, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return fmt.Errorf("Dial ControlAgent Error. %s", err.Error())
	}

	c := tio_control_v1.NewControlServiceClient(conn)

	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()

	reply, err := c.GetPloy(ctx, &tio_control_v1.TioPloy{})
	if err != nil {
		return err
	}

	for name, value := range reply.Ploy {
		m.ploy[name] = int(value)
	}

	return nil
}

func (m monImplement) UpdatePloy(name string, connect int) {
	m.ploy[name] = connect
}

func (m monImplement) GetPloy() map[string]int {
	return m.ploy
}

func (m monImplement) NoticeProxyService(name, endpoint string) error {
	for _, p := range m.proxyService {
		if err := m.invokeProxyService(p, name, endpoint); err != nil {
			logrus.Errorf("Notice proxy service error. %s", err.Error())
			continue
		}
	}

	return nil
}

func (m monImplement) invokeProxyService(add, name, endpoint string) error {
	conn, err := grpc.Dial(add, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return fmt.Errorf("Dial ProxyAgent Error. %s", err.Error())
	}

	c := tio_control_v1.NewProxyServiceClient(conn)
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()

	reply, err := c.NewEndpoint(ctx, &tio_control_v1.ProxyEndpointRequest{
		Name:     name,
		Endpoint: endpoint,
	})

	if err != nil {
		return err
	}

	if reply.Code != tio_control_v1.CommonRespCode_RespSucc {
		return fmt.Errorf("Name:  %s Endpoint: %s Error: %s", name, endpoint, err.Error())
	}

	return nil
}

// NeedScala 是否达到扩缩容标准
// 判断标准:
// 如果TrafficCount >= ploy*(2+N), 则扩容N倍。
// 如果ploy / 2 =<TrafficCount < ploy*2 , 则保持现状
// 如果0< TrafficCount < ploy/2, 则缩容1倍
// 如果 TrafficCount == 0,完成 缩容
func (m monImplement) NeedScala(traffic envoyTraffic) (bool, float64) {
	ploy, exist := m.ploy[traffic.Name]
	if exist {

		if traffic.TrafficCount == 0 {
			return true, 0
		}

		if traffic.TrafficCount >= ploy*2 {
			m := int(math.Floor(float64((traffic.TrafficCount - ploy*2) / ploy)))
			if m >= 1 {
				return true, float64(m) + 2
			}
			return true, 2
		}

		if traffic.TrafficCount > ploy/2 && traffic.TrafficCount < ploy*2 {
			return false, 0
		}

		return true, float64(1) / 2

	} else {
		return false, 0
	}

}

func (m monImplement) queryConnectInSecnodRange(query string, stepVal int) (int, error) {
	return m.proImp.QueryRange(query, StepSecond, stepVal)
}

func (m monImplement) queryConnectInMinuteRange(query string, stepVal int) (int, error) {
	return m.proImp.QueryRange(query, StepMinute, stepVal)
}
func (m monImplement) queryConnectInHourdRange(query string, stepVal int) (int, error) {
	return m.proImp.QueryRange(query, StepHour, stepVal)
}

type prometheusImplement struct {
	cli api.Client
}

func (p prometheusImplement) QueryRange(query string, step Step, stepVal int) (int, error) {
	var connectCount int

	v1api := v1.NewAPI(p.cli)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r := v1.Range{
		End: time.Now(),
	}

	switch step {
	case StepSecond:
		r.Start = time.Now().Add(-1 * time.Duration(stepVal) * time.Second)
		r.Step = time.Second
	case StepMinute:
		r.Start = time.Now().Add(-1 * time.Duration(stepVal) * time.Minute)
		r.Step = time.Minute
	case StepHour:
		r.Start = time.Now().Add(-1 * time.Duration(stepVal) * time.Hour)
		r.Step = time.Minute
	}

	query = fmt.Sprintf("sum(envoy_cluster_upstream_cx_active{envoy_cluster_name=~\"%s\"})", query)

	result, warnings, err := v1api.QueryRange(ctx, query, r)
	if err != nil {
		return connectCount, fmt.Errorf("Error querying Prometheus: %v\n", err)
	}

	if len(warnings) > 0 {
		return connectCount, fmt.Errorf("Warnings: %v\n", warnings)
	}

	m, ok := result.(model.Matrix)
	if ok {
		var connect float64
		var i int
		for _, mtrics := range m {
			for _, v := range mtrics.Values {
				connect += float64(v.Value)
				i++
			}
		}

		connectCount = int(math.Floor(connect / float64(i)))
	}

	return connectCount, nil
}

func (p prometheusImplement) QueryAllCluster() ([]string, error) {
	var clusters []string

	v1api := v1.NewAPI(p.cli)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, warnings, err := v1api.LabelValues(ctx, "envoy_cluster_name")
	if err != nil {
		return clusters, fmt.Errorf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		return clusters, fmt.Errorf("Warnings: %v\n", warnings)
	}

	for _, r := range result {
		clusters = append(clusters, string(r))
	}

	return clusters, nil
}
