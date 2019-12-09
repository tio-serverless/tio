# TIO
> 一个通过Goalng实现的Serveless平台


## Tio如何工作?

Tio 简要设计图如下:

![](https://tva1.sinaimg.cn/large/006tNbRwly1g9q84ef31rj31cx0u0n1r.jpg)

Tio分为部署态和运行态:

* `部署态`包括`Build`和`Deploy`两个阶段。 

* `运行态`指的是`Envoy`接受真实请求并转发到后端Serveless阶段。

