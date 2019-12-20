package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := Hello(context.Background(), &String{Value: "world!"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
