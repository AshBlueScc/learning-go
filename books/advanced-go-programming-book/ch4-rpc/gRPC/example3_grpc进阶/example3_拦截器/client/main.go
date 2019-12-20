package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example3_拦截器/protobuf"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto2.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto2.HelloRequest{Name: "World!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Client: %s", r.Message)
}

