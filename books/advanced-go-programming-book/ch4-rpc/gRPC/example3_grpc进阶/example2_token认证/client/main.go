package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example2_token认证/protobuf"
)

func main() {
	auth := proto2.Authentication{
		Login:    "gopher",
		Password: "password",
	}

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto2.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &proto2.HelloRequest{Name: " World!"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Client: %s", reply.Message)
}
