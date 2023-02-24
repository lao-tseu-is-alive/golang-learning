package main

import (
	"fmt"
	"os"

	"github.com/go-ping/ping"
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
	fmt.Println("This code attempts to send an unprivileged ping via UDP. On linux, this must be enabled by setting:")
	fmt.Println("sudo sysctl -w net.ipv4.ping_group_range=\"0   2147483647\"")
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run() // blocks until finished
	if err != nil {
		fmt.Println(err)
	}
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)
}
