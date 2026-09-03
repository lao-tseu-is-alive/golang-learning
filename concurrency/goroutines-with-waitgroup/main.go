package main

import (
	"fmt"
	"sync"
	"time"
)

func sayWithWG(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d \t %s\n", i, s)
	}
}

func main() {
	fmt.Printf("## Starting the main program ##\n")
	fmt.Printf("## Creating the sync.WaitGroup ##\n")
	wg := &sync.WaitGroup{}
	wg.Add(2) // Add 2 goroutines to wait for
	fmt.Printf("## Running first go routine with world ##\n")
	go sayWithWG("world", wg)
	fmt.Printf("## Running second go routine with hello ##\n")
	go sayWithWG("hello", wg)
	wg.Wait() // Wait for all goroutines to finish
}
