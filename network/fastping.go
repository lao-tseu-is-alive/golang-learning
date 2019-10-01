package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing ip argument")
		os.Exit(1)
	}
	ip := os.Args[1]
	if len(ip) < 8 {
		fmt.Println("Invalid ip argument")
		os.Exit(1)
	}
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
