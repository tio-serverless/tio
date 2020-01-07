package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/hashicorp/consul/api"
	w "github.com/hashicorp/consul/api/watch"
	"github.com/sirupsen/logrus"
)

type ActionType int
type RouteType int

const (
	GRPCRoute = iota
	HTTPRoute
	TCPRoute
)

const (
	CheckEvent = iota
	KvEvent
	AddEndpoint
	RemoveEndpoint
)

type consulData struct {
	etype   ActionType
	status  string
	content string
	name    string
}

type service struct {
	name      string
	endpoint  string
	url       string
	routetype RouteType
	action    ActionType
	remove    bool
}

type consulCli interface {
	watchKVEvents(chan consulData)
	watchChecks(chan consulData)
	getServices() map[string][]service
	handlerRoute(map[string][]service) error
	handlerKVEvent(cd consulData)
	handlerCheckEvent(cd consulData) map[string][]service
	routeInit() error
}

// routeConvert2RouterDiscoveryResponse 转换成Envoy Route
func (c client) routeConvert2RouterDiscoveryResponse(consulRoute map[string][]service) *envoy_api_v2.DiscoveryResponse {
	var envoyRoute *envoy_api_v2.DiscoveryResponse
	var routes []*route.Route

	for name, srv := range consulRoute {
		if len(srv) == 0 {
			continue
		}

		url := srv[0].url
		srvType := srv[0].routetype
		isRemove := srv[0].remove

		var r *route.Route

		defaultHttpAction := &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{Cluster: fmt.Sprintf("%s_cluster", c.defaultCluster[HTTPRoute])},
				PrefixRewrite:    "/",
			}}

		defaultGrpcction := &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{Cluster: fmt.Sprintf("%s_cluster", c.defaultCluster[GRPCRoute])},
				PrefixRewrite:    "/",
			}}

		//defaultTcpAction := &route.Route_Route{
		//	Route: &route.RouteAction{
		//		ClusterSpecifier: &route.RouteAction_Cluster{Cluster: fmt.Sprintf("%s_cluster", c.defaultCluster[TCPRoute])},
		//		PrefixRewrite:    "/",
		//	}}

		if srvType == HTTPRoute {
			r = &route.Route{
				Match: &route.RouteMatch{
					PathSpecifier: &route.RouteMatch_Prefix{Prefix: url},
				},
			}

			if !isRemove {
				r.Action = &route.Route_Route{
					Route: &route.RouteAction{
						ClusterSpecifier: &route.RouteAction_Cluster{Cluster: fmt.Sprintf("%s_cluster", name)},
						PrefixRewrite:    "/",
					}}
			} else {
				r.Action = defaultHttpAction
			}

		} else {
			r = &route.Route{
				Match: &route.RouteMatch{
					PathSpecifier: &route.RouteMatch_Prefix{Prefix: url},
				},
			}
			if !isRemove {
				r.Action = &route.Route_Route{
					Route: &route.RouteAction{
						ClusterSpecifier: &route.RouteAction_Cluster{Cluster: fmt.Sprintf("%s_cluster", name)},
					}}
			} else {
				r.Action = defaultGrpcction
			}
		}

		routes = append(routes, r)
	}

	var resource []*any.Any

	rc := []*envoy_api_v2.RouteConfiguration{
		&envoy_api_v2.RouteConfiguration{
			Name: "tio",
			VirtualHosts: []*route.VirtualHost{
				&route.VirtualHost{
					Name: "consul_service",
					Domains: []string{
						"*",
					},
					Routes: routes,
				},
			},
		},
	}

	for _, rca := range rc {
		data, err := proto.Marshal(rca)
		if err != nil {
			logrus.Errorf("Marshal Error. %s", err)
			continue
		}

		resource = append(resource, &any.Any{
			TypeUrl: "type.googleapis.com/envoy.api.v2.RouteConfiguration",
			Value:   data,
		})
	}

	envoyRoute = &envoy_api_v2.DiscoveryResponse{
		VersionInfo: "1",
		Resources:   resource,
		Canary:      false,
		TypeUrl:     "type.googleapis.com/envoy.api.v2.RouteConfiguration",
		Nonce:       time.Now().String(),
	}

	return envoyRoute
}

//routeConvert2ClusterDiscoveryResponse 将Consul的路由信息转换成Envoy Cluster
func (c client) routeConvert2ClusterDiscoveryResponse(route map[string][]service) *envoy_api_v2.DiscoveryResponse {

	var resource []*any.Any

	for name, srv := range route {
		if len(srv) == 0 {
			logrus.Debugf("%s Has No Endpoints. ", name)
			continue
		}

		srvType := srv[0].routetype

		var end []*endpoint.LbEndpoint

		for _, s := range srv {
			var ip string
			var port uint32

			es := strings.Split(s.endpoint, ":")
			if len(es) == 1 {
				port = uint32(80)
			} else {
				p, _ := strconv.Atoi(es[1])
				port = uint32(p)
			}

			ip = es[0]

			end = append(end, &endpoint.LbEndpoint{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol:      core.SocketAddress_TCP,
									Address:       ip,
									PortSpecifier: &core.SocketAddress_PortValue{PortValue: port},
								},
							},
						},
					},
				},
				HealthStatus: core.HealthStatus_UNKNOWN,
			})

		}

		r := &envoy_api_v2.Cluster{
			Name: fmt.Sprintf("%s_cluster", name),
			ConnectTimeout: &duration.Duration{
				Seconds: 5,
			},
			ClusterDiscoveryType: &envoy_api_v2.Cluster_Type{},
			LbPolicy:             envoy_api_v2.Cluster_LEAST_REQUEST,
			LoadAssignment: &envoy_api_v2.ClusterLoadAssignment{
				ClusterName: fmt.Sprintf("%s_cluster", name),
				Endpoints: []*endpoint.LocalityLbEndpoints{
					&endpoint.LocalityLbEndpoints{
						LbEndpoints: end,
					},
				},
			},
			HealthChecks: []*core.HealthCheck{
				&core.HealthCheck{
					Timeout: &duration.Duration{
						Seconds: 5,
					},
					Interval: &duration.Duration{
						Seconds: 5,
					},
					InitialJitter: &duration.Duration{
						Seconds: 5,
					},
					UnhealthyThreshold: &wrappers.UInt32Value{
						Value: 5,
					},
					HealthyThreshold: &wrappers.UInt32Value{
						Value: 1,
					},
					HealthChecker: &core.HealthCheck_TcpHealthCheck_{
						TcpHealthCheck: &core.HealthCheck_TcpHealthCheck{
						},
					},
				},
			},
		}
		if srvType == GRPCRoute {
			r.Http2ProtocolOptions = &core.Http2ProtocolOptions{}
		}

		rc := []*envoy_api_v2.Cluster{
			r,
		}

		for _, rca := range rc {
			data, err := proto.Marshal(rca)
			if err != nil {
				logrus.Errorf("Marshal Error. %s", err)
				continue
			}

			resource = append(resource, &any.Any{
				TypeUrl: "type.googleapis.com/envoy.api.v2.Cluster",
				Value:   data,
			})
		}

	}

	return &envoy_api_v2.DiscoveryResponse{
		VersionInfo: "1",
		Resources:   resource,
		Canary:      false,
		TypeUrl:     "type.googleapis.com/envoy.api.v2.Cluster",
		Nonce:       time.Now().String(),
	}
}

// routeInit 初始化当前可用的路由信息
func (c client) routeInit() error {
	val, _, err := c.cli.KV().List("tio/v1/gateway/services", nil)
	if err != nil {
		return err
	}

	for _, k := range val {
		var m meta
		if err := json.Unmarshal(k.Value, &m); err == nil {
			name := trimKey(k.Key)
			ses, _, err := c.cli.Health().Service(name, "", true, nil)
			if err != nil {
				logrus.Errorf("Query [%s] Endpoints Error. %s", name, err.Error())
				continue
			}

			var srvs []service
			for _, s := range ses {
				srvs = append(srvs, service{
					name:      name,
					endpoint:  fmt.Sprintf("%s:%d", s.Service.Address, s.Service.Port),
					url:       m.Url,
					routetype: RouteType(m.RouteType),
					action:    AddEndpoint,
					remove:    m.Remove,
				})
			}

			if len(srvs) > 0 {
				c.routes[name] = srvs
			}
		}
	}

	return nil
}

func (c client) clusterInit() {
	if os.Getenv("TIO_CONSUL_CLUSTER_HTTP") != "" {
		c.defaultCluster[HTTPRoute] = os.Getenv("TIO_CONSUL_CLUSTER_HTTP")
	}

	if os.Getenv("TIO_CONSUL_CLUSTER_GRPC") != "" {
		c.defaultCluster[GRPCRoute] = os.Getenv("TIO_CONSUL_CLUSTER_GRPC")
	}

	if os.Getenv("TIO_CONSUL_CLUSTER_TCP") != "" {
		c.defaultCluster[TCPRoute] = os.Getenv("TIO_CONSUL_CLUSTER_TCP")
	}
}

func watch(cli consulCli, cc chan consulData) {
	for {
		select {
		case c := <-cc:
			switch c.etype {
			case CheckEvent:
				s := cli.handlerCheckEvent(c)
				if s != nil {
					if err := cli.handlerRoute(s); err != nil {
						logrus.Errorf("Handler Route Error: %s", err.Error())
					}
				}
			case KvEvent:
				cli.handlerKVEvent(c)
			}
		}

	}
}

func (c client) handlerCheckEvent(cd consulData) map[string][]service {
	//route := make(map[string][]service)

	if cd.name != "" {
		alive, err := c.queryAliveService(cd.name)
		if err != nil {
			logrus.Errorf("query alive service error: %s", err)
			return nil
		}

		m := c.route[cd.name]
		if m.Url == "" {
			logrus.Errorf("Can not find URL for %s", cd.name)
			return nil
		}

		var s []service

		for _, a := range alive {
			s = append(s, service{
				endpoint:  a,
				url:       m.Url,
				routetype: RouteType(m.RouteType),
				action:    AddEndpoint,
				remove:    m.Remove,
			})
		}

		c.routes[cd.name] = s
		return c.routes
	}

	return nil
}

func (c client) handlerKVEvent(cd consulData) {
	var m meta
	if err := json.Unmarshal([]byte(cd.content), &m); err == nil {
		c.route[cd.name] = m
	}
}

type client struct {
	route          map[string]meta
	cli            *api.Client
	cc             chan consulData
	address        string
	routes         map[string][]service
	defaultCluster map[int]string
}

func (c client) watchKVEvents(cd chan consulData) {
	params := make(map[string]interface{})
	params["prefix"] = "tio/v1/gateway/services"
	params["type"] = "keyprefix"

	p, err := w.Parse(params)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	p.HybridHandler = func(val w.BlockingParamVal, i interface{}) {
		if v, ok := i.(api.KVPairs); ok {
			for _, kv := range v {

				cd <- consulData{
					etype:   KvEvent,
					content: string(kv.Value),
					name:    trimKey(kv.Key),
				}

			}
		}
	}

	if err = p.Run(c.address); err != nil {
		logrus.Fatal(err.Error())
	}
}

func (c client) watchChecks(cd chan consulData) {
	params := make(map[string]interface{})
	params["type"] = "checks"

	p, err := w.Parse(params)
	if err != nil {
		log.Fatal(err)
	}

	p.HybridHandler = func(val w.BlockingParamVal, i interface{}) {
		if v, ok := i.([]*api.HealthCheck); ok {
			for _, v := range v {
				cd <- consulData{
					etype:  CheckEvent,
					status: v.Status,
					name:   v.ServiceName,
				}
			}
		}
	}

	if err = p.Run(c.address); err != nil {
		panic(err)
	}
}

func (c client) getServices() map[string][]service {
	return nil
}

func (c client) handlerRoute(endpoints map[string][]service) error {
	logrus.Debugf("route: %v", endpoints)

	c.routes = endpoints
	send2Envoy(&c)

	return nil
}

func (c client) queryAliveService(sid string) ([]string, error) {
	allServices, _, err := c.cli.Catalog().Service(sid, "", nil)
	if err != nil {
		return nil, err
	}

	var endpoints []string

	for _, s := range allServices {
		if s.Checks.AggregatedStatus() == "passing" {
			endpoints = append(endpoints, fmt.Sprintf("%s:%d", s.ServiceAddress, s.ServicePort))
		}
	}

	return endpoints, nil
}

func initClient() (*client, error) {

	config := api.DefaultConfig()
	config.Address = strings.Split(os.Getenv("CONSUL_ADDRESS"), ";")[0]

	cli, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &client{
		cli:            cli,
		route:          make(map[string]meta),
		cc:             make(chan consulData, 100),
		address:        config.Address,
		routes:         make(map[string][]service),
		defaultCluster: make(map[int]string),
	}, nil
}

func trimKey(key string) string {
	ks := strings.Split(key, "/")

	if len(ks) == 0 {
		return key
	}

	return ks[len(ks)-1]
}

func decodeValue(v []byte) (meta, error) {
	var m meta

	err := json.Unmarshal(v, &m)
	return m, err
}
