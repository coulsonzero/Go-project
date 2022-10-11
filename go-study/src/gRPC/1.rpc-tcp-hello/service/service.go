package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatalf("rpc service regist: %s \n", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
