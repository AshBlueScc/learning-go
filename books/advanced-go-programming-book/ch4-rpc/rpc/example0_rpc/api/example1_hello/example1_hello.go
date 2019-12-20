package api

import (
	"net/rpc"
)

//更安全的rpc接口
//拆分三种开发角色：
//1.服务端实现RPC方法的开发人员
//2.客户端调用RPC方法的人员
//3.制定服务端和客户端RPC接口规范的设计人员

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

//rpc服务的接口规范分为三个部分：
//1.服务的名字（这个是RPC服务抽象的包路径，并非完全等价Go语言的包路径）。移到公共部分const里面
//2.服务要实现的详细方法列表
//3.注册该类型服务的函数

//对客户端的封装
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
