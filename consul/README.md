# Consul Service

### SideCar

### Watcher
> 监控Consul KV变化，生成Envoy数据并推送给Envoy

+ Environment
  
   - CONSUL_ADDRESS
   - DEBUG
   - TIO_CONSUL_CLUSTER_HTTP
   - TIO_CONSUL_CLUSTER_GRPC
   - TIO_CONSUL_CLUSTER_TCP
   - MY_GRPC_PORT