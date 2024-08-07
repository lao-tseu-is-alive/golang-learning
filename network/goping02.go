package main

import (
	"fmt"
	"os"
	"os/signal"

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
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run() // blocks until finished
	if err != nil {
		fmt.Println(err)
	}
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)
}
