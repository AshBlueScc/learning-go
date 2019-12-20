package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example1_CA认证/protobuf"
)

var (
	port = ":5000"

	tlsDir        = "gRPC/example3_grpc进阶/example1_CA认证/keyAndCrt"

	ca0       = tlsDir + "/ca.crt"
	serverCrt = tlsDir + "/grpcService.crt"
	serverKey = tlsDir + "/grpcService.key"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *proto2.String, ) (*proto2.String, error) {
	reply := &proto2.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	certificate, err := tls.LoadX509KeyPair(serverCrt, serverKey)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca0)
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(creds))

	proto2.RegisterHelloServiceServer(server, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	_ = server.Serve(lis)
}
