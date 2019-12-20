package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type HelloServiceImpl struct {}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String, ) (*String, error) {
	reply := &String{Value: "[Hello]hello:" + args.GetValue()}
	return reply, nil
}


func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &String{Value: "[Channel]hello:" + args.GetValue()}

		err = Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}