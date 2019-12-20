package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example4_grpc和protobuf扩展/example1_rest接口/proto"
)

type myGrpcServer struct{}

func (s *myGrpcServer) Get(ctx context.Context, in *proto2.StringMessage) (*proto2.StringMessage, error) {
	return &proto2.StringMessage{Value: "Get: " + in.Value}, nil
}

func (s *myGrpcServer) Post(ctx context.Context, in *proto2.StringMessage) (*proto2.StringMessage, error) {
	return &proto2.StringMessage{Value: "Post: " + in.Value}, nil
}

func main() {
	server := grpc.NewServer()
	proto2.RegisterRestServiceServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Panicf("could not list on %s:%s", ":5000", err)
	}

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}