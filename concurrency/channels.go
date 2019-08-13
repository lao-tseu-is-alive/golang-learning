package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

/*
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
The data flows in the direction of the arrow.
*/

func main() {
	s := []int{7, 3, 8, -9, 2, 0, -1, 100}
	// Like maps and slices, channels must be created before use
	c := make(chan int)
	// will sums the numbers in a slice, distributing the work between two goroutines.
	// Once both goroutines have completed their computation, it calculates the final result.
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
