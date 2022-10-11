package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// $ curl localhost:1234/jsonrpc -X POST \
// --data '{"method":"HelloService.Hello","params":["world"],"id":0}'
// {"id":0,"result":"hello:world","error":null}

func main() {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatalf("rpc service regist: %s \n", err)
	}

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		conn := struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	log.Fatal(http.ListenAndServe(":1234", nil))
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
