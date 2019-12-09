# Consul Agent
> 负责服务注册 


`Consul-Agent` 会作为sidecar伴随service同时部署到K8s集群中，然后将服务信息注册到Consul。 使用时需要提供以下环境变量:

* MY_POD_NAME Pod名称，将会作为服务实例注册到Consul中
* MY_POD_IP PodIP
* MY_POD_PORT 服务端口。用作Consul Health Check中的一部分
* MY_SERVICE_NAME 服务名称，同一组Deployment需要保持一致(如果是Grpc服务，此属性应该是Grpc服务名称)
* CONSUL_ADDRESS consul server地址，使用';'分割

其中`Pod_Name`和`Pod_IP`建议采取注入方式设置，如下：

```yaml
        - name: MY_POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: MY_POD_IP
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: status.podIP
```

