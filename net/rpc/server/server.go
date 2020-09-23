package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// 参数结构体
type Args struct {
	 A int
	 B int
}

//
type Quotient struct {
	Quo, Rem int
}

//
type Arith int

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, quo *Quotient) error  {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func Run()  {
	arith := new(Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatalln("rpc.Register() error:", err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln("net.Listen error:", err)
	}
	http.Serve(l, nil)
}

