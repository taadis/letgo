package main

import (
	"io"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
}

func main() {
	ln, err := net.Listen("tcp", ":5903")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("echo-server listening on %s", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("accepted connection to %s from %s", conn.LocalAddr().String(), conn.RemoteAddr().String())

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			log.Println("client closed")
			break
		}
		if err != nil {
			log.Println(err)
			return
		}

		conn.Write(buf[:n])
	}
}
