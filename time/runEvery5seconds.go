package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c)
	start := time.Now()
	fmt.Println("Starting Timer !")
	ticker := time.NewTicker(5 * time.Second) // every 5 second
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				elapsed := time.Since(start)
				fmt.Printf("Tick (after %s)\n", elapsed.Round(time.Millisecond))
			case <-stop:
				fmt.Println("Goroutine closing")
				return
			}
		}
	}()

	// Block until
	// the signal is received
	<-c
	ticker.Stop()

	// Stop the goroutine
	stop <- true
	// Wait until the
	<-stop
	fmt.Println("Application stopped")
}
