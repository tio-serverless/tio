package main

type monitorInterface interface {
	// WatchProemetheus 监控Prometheus实时流量
	WatchProemetheus() (chan []envoyTraffic, error)
	// Sacla 扩缩容
	Sacla(name string, num float64) error
	// WaitScala 等待服务扩缩容结束
	WaitScala(name string) (string, error)
	//// IsScalaSucc 扩缩容是否成功
	//IsScalaSucc(name string) (bool, error)
	// InvokeDeployService 调用部署服务
	InvokeDeployService(name string, num float64) error
	// InitPloy 初始化扩容策略
	InitPloy() error
	// UpdatePloy 更新扩容策略
	UpdatePloy(string, int)
	// GetPloy 获取最新策略
	GetPloy() map[string]int
	// NoticeProxyService 通知反向代理服务
	NoticeProxyService(name, endpoint string) error
	// NeedScala 是否达到扩缩容标准
	NeedScala(Traffic envoyTraffic) (bool, float64)
}

type prometheusInterface interface {
	// QueryRange 查询一段时间内某个指标的平均值
	QueryRange(query string, step Step, stepVal int) (int, error)
	QueryAllCluster() ([]string, error)
}

type envoyTraffic struct {
	Name         string
	TrafficCount int
}

const (
	StepSecond Step = iota
	StepMinute
	StepHour
)

type Step int
