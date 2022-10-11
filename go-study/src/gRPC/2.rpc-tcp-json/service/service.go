package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 使用nc命令行工具实现rpc方法调用
//  $ nc -l 1234
// 启动service和client终端将出现：
// {"method":"HelloService.Hello","params":["world"],"id":0}

func main() {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatalf("rpc service regist: %s \n", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
