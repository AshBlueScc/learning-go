package main

import (
	"context"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
	"xorm/books/advanced-go-programming-book/ch4-rpc/gRPC/example2_grpc发布和订阅模式/pubsubservice"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *pubsubservice.String, ) (*pubsubservice.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pubsubservice.String{}, nil
}

func (p *PubsubService) Subscribe(arg *pubsubservice.String, stream pubsubservice.PubsubService_SubscribeServer, ) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pubsubservice.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	pubsubservice.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
