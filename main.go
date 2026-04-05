package main

import (
	"errors"
	"flag"
	"log"

	am "github.com/adaptivemsg/adaptivemsg-go"
	hello "github.com/adaptivemsg/hello-server-rust/api/hello"
	//hello "github.com/adaptivemsg/hello-server-go/api/hello"
)

func main() {
	addr := flag.String("addr", "tcp://127.0.0.1:5555", "server address (examples: tcp://127.0.0.1:5555, uds://@adaptivemsg-hello, uds:///tmp/adaptivemsg-hello.sock)")
	flag.Parse()

	conn, err := am.NewClient().Connect(*addr)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := am.SendRecvAs[*hello.HelloReply](conn, &hello.HelloRequest{
		Who:      "John",
		Question: "who are you",
	})
	if err != nil {
		log.Printf("reply error: %v", err)
		return
	}
	log.Printf("reply: %s (trace %s)", reply.Answer, reply.Internal.TraceID)

	_, err = am.SendRecvAs[*hello.HelloReply](conn, &hello.HelloRequest{
		Who:      "Bob",
		Question: "error please",
	})
	if err != nil {
		var remote am.ErrRemote
		if errors.As(err, &remote) {
			log.Printf("expected error: %s: %s", remote.Code, remote.Message)
			return
		}
		log.Printf("unexpected error: %v", err)
		return
	}
	log.Printf("unexpected success")
}
