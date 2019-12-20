package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example3_拦截器/protobuf"
)

var(
	port = ":1234"
)

type myGrpcServr struct {}

func (s *myGrpcServr) SayHello(ctx context.Context, in *proto2.HelloRequest)(*proto2.HelloReply, error){
	//如果报错就用defer func处理
	//panic("debug")
	//如果没报错就用handler处理
	return &proto2.HelloReply{Message: "Hello " + in.Name}, nil
}

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(resp interface{}, err error){
	log.Println("filter: ", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))
	proto2.RegisterGreeterServer(server, new(myGrpcServr))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}