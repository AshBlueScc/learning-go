package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"xorm/books/advanced-go-programming-book/ch4-rpc/rpc/example0_rpc/server/example0_service_rpc入门/service"
)

func main() {
	rpc.RegisterName("HelloService", new(example0.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}


//还有一部分用nc命令调用的，暂时没有下载虚拟机，要用linux系统
//以后测试