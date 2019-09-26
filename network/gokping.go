package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/kping"
	"time"
)

/*
	requirements :
	sudo apt-get install libpcap-dev
	go get -u github.com/lao-tseu-is-alive/kping

*/

func main() {
	// Create a new Pinger
	pinger, err := kping.NewPinger("192.168.50.7", 100, 10, 1*time.Minute, 100*time.Millisecond)
	if err != nil {
		golog.Err("creating kping.NewPinger : %v", err)
	}

	// Add IP addresses to Pinger
	if err := pinger.AddIPs([]string{"192.168.50.4", "8.8.8.8"}); err != nil {
		golog.Err("adding ips adr to kping.NewPinger : %v", err)
	}

	// Run !
	statistics, err := pinger.Run()
	if err != nil {
		golog.Err("doing  pinger.Run : %v", err)
	}

	// Print result
	for ip, statistic := range statistics {
		fmt.Printf("%s: %v\n", ip, statistic)
	}
}
