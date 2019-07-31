package main

import (
	"fmt"
	"net"
)

func main() {

	// Get all network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range ifaces {
		// Resolve addresses
		// for each interface
		addres, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Println(i.Name)
		for _, add := range addres {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}

	}
}
