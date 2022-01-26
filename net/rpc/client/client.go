package client

import (
	"log"
	"net/rpc"

	"github.com/taadis/letgo/net/rpc/server"
)

func Run() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 同步调用
	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Println("client.Call Arith.Multiply error:", err)
	}
	log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
	log.Println()
}
