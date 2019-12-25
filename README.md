# tio
 
[![Build Status](https://travis-ci.org/tio-serverless/tio.svg?branch=master)](https://travis-ci.org/tio-serverless/tio)

Golang Serverless Framework

## Tio是什么？

Tio是一个支持golang的Serverless平台，目前支持也仅支持golang。 

## Tio支持哪些运行模式?

Tio当前支持以下三种运行模式:

- Http Event
- Grpc Event
- AMQP Event

## 如何在Tio部署Function?

完成一次Function部署，需要以下三步：

1. 代码实现
2. 完成配置
3. 上传部署

    + 代码实现
    
    每种模式有一种Code Template， 默认情况下，仅需要使用你的业务逻辑**填充**里面空白部分即可。 
    ```golang
    # Grpc 
    package main
    
    import (
        "context"
    
        "google.golang.org/grpc"
    )
    
    // register 
    // Register your grpc server.
    func register(s *grpc.Server, srv *server) {
        // Please invoke your grpc register funcion, e.g. rpc.RegisterEchoServer(s, srv)
        
    }
    
    
    // type server struct{}
    //
    // Server as the truly grpc server instance, it has been declared.
    // So please implement GRPC function as the blowing:
    // func (s server) Hello(context.Context, *rpc.HelloRequest) (*rpc.HelloResponse, error) {
    //		return &rpc.HelloResponse{}, nil
    //	}
    //
    // If you want to initialize server struct, please type code in the flowing function.
    //
    // func (s server) ServerInit(){
    //	 panic("Please Implement Me!")
    // }
    ```
    
    + 完成配置
    
    在代码同级目录中，需要存在`.tio.toml`文件。 此文件描述了Tio应该如何部署当前代码。
    
    + 上传部署
    
    执行`tio-cli deploy -d .` 就会将当前您所编写的代码打包上传，进行实例部署。
    
