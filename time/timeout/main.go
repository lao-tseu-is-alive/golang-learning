package main

import (
	"fmt"
	"time"
)

func main() {

	timeout := 5
	start := time.Now()
	to := time.After(time.Duration(timeout) * time.Second)
	list := make([]string, 0)
	done := make(chan bool, 1)

	fmt.Printf("Starting to insert items with a timeout (%d seconds)\n", timeout)
	go func() {
		defer fmt.Println("Exiting goroutine")
		for {
			select {
			case <-to:
				elapsed := time.Since(start)
				fmt.Printf("The time is up after %s !\n", elapsed.Round(time.Millisecond))
				done <- true
				return
			default:
				list = append(list, time.Now().String())
			}
		}
	}()

	<-done
	fmt.Printf("Managed to insert %d items\n", len(list))
}
