package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type HelloServer string

func (s HelloServer) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	rw.Write([]byte(string(s)))
}

func createMyHttpServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: HelloServer("HELLO GOPHER!\n"),
	}
}

const defaultHttpAddr = "localhost:7070"

func main() {
	s := createMyHttpServer(defaultHttpAddr)
	go s.ListenAndServe()

	// Connect with plain TCP
	conn, err := net.Dial("tcp", defaultHttpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost:localhost:7070\r\n\r\n")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(time.Second))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	ctx, _ := context.WithTimeout(context.Background(),
		5*time.Second)
	s.Shutdown(ctx)

}
