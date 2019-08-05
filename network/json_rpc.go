package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Div(args *Args, reply *Reply) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.C = args.A / args.B
	return nil
}

func startServer() {
	arith := new(Arith)

	server := rpc.NewServer()
	server.Register(arith)

	l, e := net.Listen("tcp", ":8222")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func main() {

	// starting server in go routine (it ends on end
	// of main function
	go startServer()

	// now client part connecting to RPC service
	// and calling methods

	conn, err := net.Dial("tcp", "localhost:8222")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := jsonrpc.NewClient(conn)

	var reply Reply
	var args *Args
	for i := 0; i < 11; i++ {
		args = &Args{7, i}
		err = c.Call("Arith.Div", args, &reply)
		if err != nil {
			fmt.Printf("arith error: %s\n", err.Error())
		} else {
			fmt.Printf("Arith: %d / %d = %v\n", args.A, args.B, reply.C)
		}

		err = c.Call("Arith.Add", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d + %d = %v\n", args.A, args.B, reply.C)

		// NL
		fmt.Printf("\033[33m%s\033[m\n", "---------------")

	}
}
