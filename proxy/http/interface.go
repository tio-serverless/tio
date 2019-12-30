package main

import "net/http"

type dataLoader interface {
	//loadInjectData 加载Inject数据,并完成初始化
	LoadInjectData() error
	//scala 通知monitor service 进行扩容
	Scala(string) error
	//wait 等待指定的服务实例创建完成
	Wait(string) (service, error)
	//done 指定的服务创建成功
	Done(service) error
	////proxy 反向代理外部请求
	//Proxy(http.ResponseWriter, *http.Request)
	//isValidInject 判断传入的服务是否是已经被缓存
	IsValidInject(string) bool
	//getServiceName 使用URL换取相对应的服务名称
	GetServiceName(string) string
	//Transfer 透明传输
	Transfer(string, http.ResponseWriter, *http.Request)
}

type service struct {
	Name     string
	Endpoint string
}
