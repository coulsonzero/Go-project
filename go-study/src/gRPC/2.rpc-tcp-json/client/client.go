package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// client, err := rpc.Dial("tcp", "localhost:1234")
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("rpc client dial: %s \n", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	if err = client.Call("HelloService.Hello", "world", &reply); err != nil {
		log.Fatalf("rpc client call: %s \n", err)
	}
	fmt.Println(reply)
}
