package main

import (
	"fmt"
	"net"
)

func resolveByIp(ip string) {
	// Resolve by IP
	addrs, err := net.LookupAddr(ip)
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Printf("[%s]\t%s\n", ip, addr)
	}
}

func getIpByName(hostName string) {
	//Resolve by address
	ips, err := net.LookupIP(hostName)
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Printf("%s : [%s]\n", hostName, ip.String())
	}
}

func main() {

	resolveByIp("127.0.0.1")
	resolveByIp("54.38.110.125")

	getIpByName("localhost")
	getIpByName("go.trouvl.info")

}
