package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example4_grpc和protobuf扩展/example1_rest接口/proto"
)

var (
	//port         = ":1234"
	//echoEndpoint = flag.String("echo_endpoint", "localhost"+port, "endpoint of YourService")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto2.RegisterRestServiceHandlerFromEndpoint(ctx, mux, "localhost:5000", opts)
	if err != nil {
		return
	}

	_ = http.ListenAndServe(":8080", mux)
}
