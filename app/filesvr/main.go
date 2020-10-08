package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	fmt.Println("listening and serve on", l.Addr().String())
	err = http.Serve(l, mux)
	if err != nil {
		panic(err)
	}
}
