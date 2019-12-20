package main

import (
	"fmt"
	"log"
	"xorm/books/advanced-go-programming-book/ch4-rpc/rpc/example0_rpc/api/example1_hello"
)

func main(){
	//client0_publish, err := example0_rpc.Dial("tcp", "localhost:1234")
	client, err := api.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	//err = client0_publish.Call(service.HelloServiceName+".hello", "hello", &reply)
	err = client.Hello("World!", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
