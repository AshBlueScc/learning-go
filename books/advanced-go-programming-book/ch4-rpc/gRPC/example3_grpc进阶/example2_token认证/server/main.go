package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example2_token认证/protobuf"
)

var(
	port = ":1234"

)

type myGrpcServer struct {}

func (s *myGrpcServer) SayHello(ctx context.Context, in *proto2.HelloRequest)(*proto2.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil,fmt.Errorf("missing credentials")
	}

	var (
		appid string
		appky string
	)

	if val, ok := md["login"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appky = val[0]
	}

	if appid != "gopher" || appky != "password" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token: appid=%s, appky=%s", appid, appky)
	}

	return &proto2.HelloReply{Message: "Hello" + in.Name}, nil

}

func main(){
	server := grpc.NewServer()
	proto2.RegisterGreeterServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}
