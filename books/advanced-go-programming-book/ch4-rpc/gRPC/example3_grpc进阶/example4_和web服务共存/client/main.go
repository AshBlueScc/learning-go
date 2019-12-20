package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example4_和web服务共存/protobuf"
)

var(
	port     = ":1234"
	basePath = "gRPC/example3_grpc进阶/example4_和web服务共存/tls-config/"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(basePath+"grpcService.crt", "grpcService.grpc.io")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto2.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto2.HelloRequest{Name: "World!!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Client: "+ r.Message)
}