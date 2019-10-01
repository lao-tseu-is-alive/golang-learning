package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"os"
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
