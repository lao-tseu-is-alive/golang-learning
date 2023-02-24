package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("Waiting for client...")
		fmt.Println("you can try to just type : \"nc localhost 8080\" from another shell")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Connection from %v\n", conn.RemoteAddr())

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("received empty string, conn:  %v\n", conn)
			_, err = io.WriteString(conn, "An empty string is not a good way to talk together, don't you think so ? ")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Will send back this received msg:  %v\n", msg)
			_, err = io.WriteString(conn, "Received: "+string(msg))
			if err != nil {
				fmt.Println(err)
			}
		}
		conn.Close()
	}
}
