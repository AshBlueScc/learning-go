package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	proto2 "xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example3_grpc进阶/example1_CA认证/protobuf"
)

var (
	port = ":5000"

	tlsDir        = "gRPC/example3_grpc进阶/example1_CA认证/keyAndCrt"
	tlsServerName = "grpcService.io"

	ca0 = tlsDir + "/ca.crt"
	clientCrt  = tlsDir + "/webService.crt"
	clientKey  = tlsDir + "/webService.key"
)

func main() {
	certificate, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca0)
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: false,
		Certificates:       []tls.Certificate{certificate},
		RootCAs:            certPool,
		ServerName:         tlsServerName,
	})

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto2.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto2.String{Value: "world!"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
