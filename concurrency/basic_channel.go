package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("mouton", c)

	// using range there is no need to check if chanel is open or closed
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
