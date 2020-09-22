package client

import (
	"log"
	"net/rpc"
)

func main(){
	client, err := rpc.DialHTTP("tcp",   ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
}
