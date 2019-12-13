package main

import (
	"fmt"
	"log"
	"xorm/rpc/rpc/api/example1_hello"
)

func main(){
	//client, err := rpc.Dial("tcp", "localhost:1234")
	client, err := api.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	//err = client.Call(service.HelloServiceName+".hello", "hello", &reply)
	err = client.Hello("World!", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
