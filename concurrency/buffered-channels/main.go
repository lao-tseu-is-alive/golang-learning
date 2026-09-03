package main

import "fmt"

/*
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
*/
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // free one place in buffer
	ch <- 3           // so this one is ok
	fmt.Println(<-ch) // free one place in buffer
	fmt.Println(<-ch) // free another place in buffer
	ch <- 4           // so this one is ok
	ch <- 5           // and this one also
	ch <- 6           // badaboumm overfilling the buffer => fatal error: all goroutines are asleep - deadlock!

}
