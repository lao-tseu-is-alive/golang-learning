package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		msg := strings.TrimSpace(string(netData))
		if len(msg) > 0 {
			fmt.Printf("For adr : [%v] echoing back this msg:-> %v\n", c.RemoteAddr(), msg)
			c.Write([]byte(netData))
			if strings.ToUpper(strings.TrimSpace(string(netData))) == "STOP" {
				msg = "You send me STOP, so I will close the connection.\n"
				fmt.Printf("For adr : [%v] STOP received, sending msg:-> %v\n", c.RemoteAddr(), msg)
				c.Write([]byte(msg))
				c.Close()
				break
			}
		} else {
			msg = "An empty string is not a good way to talk together, don't you think so ?\n"
			fmt.Printf("For adr : [%v] empty string received, sending msg:-> %v\n", c.RemoteAddr(), msg)
			c.Write([]byte(msg))
		}
	}
	time.Sleep(3 * time.Second)
	c.Close()
}

func main() {

	arguments := os.Args
	PORT := ":8080"
	if len(arguments) == 1 {
		fmt.Println("No port number provided as parameter using 8080 as a default!")
	} else {
		PORT = ":" + arguments[1]
	}
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		go handleConnection(c)
	}
}
