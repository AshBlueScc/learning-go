package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example2_grpc发布和订阅模式/pubsubservice"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pubsubservice.NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(),&pubsubservice.String{Value: "golang: hello Go!"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &pubsubservice.String{Value: "docker: hello Docker!"})
	if err != nil {
		log.Fatal(err)
	}
}
