package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	PORT := ":8080"
	if len(arguments) == 1 {
		fmt.Println("No port number provided as parameter using 8080 as a default!")
	} else {
		PORT = ":" + arguments[1]
	}

	s, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Printf("ERROR doing net.ResolveUDPAddr err : %v", err)
		os.Exit(100)
	}

	connection, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Printf("ERROR doing net.ListenUDP(\"udp\", %s) err : %v", s, err)
		os.Exit(100)
	}

	defer connection.Close()

	// wandering what happens if i receive more data then the size of the buffer ?!
	buffer := make([]byte, 2048)
	fmt.Printf("you can try to just type : \"nc -v -u localhost %s\" from another shell\n", PORT)
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("ERROR doing connection.ReadFromUDP(buffer) : %v\n", err)
			os.Exit(100)
		}
		fmt.Printf("received (%v)-> [%s] from %v\n", n, strings.TrimSpace(string(buffer[0:n])), addr)
		data := []byte(buffer[0:n])
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP server!")
			_, err = connection.WriteToUDP([]byte("Exiting UDP server!"), addr)
			if err != nil {
				fmt.Println(err)
				os.Exit(100)
			}
			return
		}
	}
}

/*
pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	buffer := make([]byte, 2048)
	fmt.Println("Waiting for client...")

	for {
		_, addr, err := pc.ReadFrom(buffer)
		if err == nil {
			rcvMsq := string(buffer)
			fmt.Println("Received: " + rcvMsq)
			if _, err := pc.WriteTo([]byte("Received: "+rcvMsq), addr);
				err != nil {
				fmt.Println("error on write: " + err.Error())
			}
		} else {
			fmt.Println("error: " + err.Error())
		}
	}
}*/
