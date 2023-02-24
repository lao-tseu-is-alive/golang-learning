package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Result int

type RpcServer struct{}

func (t RpcServer) Add(args *Args, result *Result) error {
	log.Printf("Adding %d to %d\n", args.A, args.B)
	*result = Result(args.A + args.B)
	return nil
}

const localAddr = ":8222"

func main() {
	go createMyServer(localAddr)
	client, err := jsonrpc.Dial("tcp", fmt.Sprintf("localhost%s", localAddr))
	if err != nil {
		golog.Err("ERROR trying jsonrpc.Dial %v", err)
		panic(err)
	}
	defer client.Close()
	args := &Args{
		A: 2,
		B: 3,
	}
	var result Result
	err = client.Call("RpcServer.Add", args, &result)
	if err != nil {
		golog.Err("error in RpcServer", err)
		os.Exit(1)
	}
	log.Printf("%d+%d=%d\n", args.A, args.B, result)
}

func createMyServer(addr string) {
	server := rpc.NewServer()
	err := server.Register(RpcServer{})
	if err != nil {
		panic(err)
	}
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalf("Couldn't start listening on %s errors: %s",
			addr, e)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
