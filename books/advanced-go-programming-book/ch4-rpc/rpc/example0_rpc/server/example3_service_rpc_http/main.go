package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"xorm/books/advanced-go-programming-book/ch4-rpc/rpc/example0_rpc/server/example0_service_rpc入门/service"
)

func main() {
	rpc.RegisterName("HelloService", new(example0.HelloService))

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer: writer,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}

//构造post请求：http://localhost:1234/jsonrpc
//参数：{"method":"HelloService.Hello","params":["hello"],"id":0}
//返回：{"id":0,"result":"hello world!","error":null}