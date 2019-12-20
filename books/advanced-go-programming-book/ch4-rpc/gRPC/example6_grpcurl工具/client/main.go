package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example6_grpcurl工具/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "world!"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
