# Inject

获取当前部署的服务元数据。并及时告之`Proxy`模块。

### Grpc

通过`grpcurl`获取当前所有的Grpc服务名称和导出的Methods，并持久化到Redis中。