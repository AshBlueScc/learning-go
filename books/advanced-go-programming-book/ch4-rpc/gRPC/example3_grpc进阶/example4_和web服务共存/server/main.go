package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"strings"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example4_和web服务共存/protobuf"
)

var (
	port     = ":1234"
	basePath = "gRPC/example3_grpc进阶/example4_和web服务共存/tls-config/"
)

type myGrpcServer struct{}

func (s *myGrpcServer) SayHello(ctx context.Context, in *proto2.HelloRequest) (*proto2.HelloReply, error) {
	return &proto2.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile(basePath+"grpcService.crt", basePath+"grpcService.key")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto2.RegisterGreeterServer(grpcServer, new(myGrpcServer))


	//如下部分不知道怎么访问，postman没有添加证书和服务器名然后访问的
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World!")
	})

	http.ListenAndServeTLS(port, basePath+"grpcService.crt", basePath+"grpcService.key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks
		// https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	}))
}
