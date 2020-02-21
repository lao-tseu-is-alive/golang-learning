package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted // to limit the number of concurrent open files in the goroutines
}

const verbose = false

func ScanPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			fmt.Printf("#NOTICE: too many open files occured, sleeping %v\n", timeout)
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			if verbose {
				fmt.Println(err.Error())
				fmt.Printf("%d \tclose\n", port)
			}
		}
		return false
	}
	conn.Close()
	fmt.Printf("%d \topen\n", port)
	return true
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(port)
	}
}

/**
based on this article :
https://medium.com/@KentGruber/building-a-high-performance-port-scanner-with-golang-9976181ec39d
I did remove the ulimit stuff from the original  because it was just used to get back
the maximum number of open files  obtained from ulimit -n = 1024
i prefer to hard code it to a reasonable value like 512
 the semaphore is here to limit the number of simultaneous opened files in  goroutines.
https://godoc.org/golang.org/x/sync/semaphore#example-package--WorkerPool
**/

func main() {
	var ipAdr string
	flag.StringVar(&ipAdr, "ip", "127.0.0.1", "ip address to scan")
	flag.Parse()
	fmt.Printf("# ports listening on ip: %s\n", ipAdr)
	fmt.Printf("# runtime.GOMAXPROCS(0) : %v\n", runtime.GOMAXPROCS(0))
	ps := &PortScanner{
		ip:   ipAdr,
		lock: semaphore.NewWeighted(512),
	}
	ps.Start(1, 65535, 50*time.Millisecond)
}
