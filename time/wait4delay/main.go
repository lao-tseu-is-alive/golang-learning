package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	t := time.NewTimer(4 * time.Second)

	fmt.Printf("Will start waiting 4 sec  at : %v\n",
		time.Now().Format(time.UnixDate))
	<-t.C
	fmt.Printf("Code executed ~4sec later at : %v\n",
		time.Now().Format(time.UnixDate))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Printf("Start waiting 7 seconds for AfterFunc at : %v\n",
		time.Now().Format(time.UnixDate))
	time.AfterFunc(7*time.Second, func() {
		fmt.Printf("Code in AfterFunc() executed 7 seconds later at : %v\n",
			time.Now().Format(time.UnixDate))
		wg.Done()
	})

	wg.Wait()

	fmt.Printf("Waiting on time.After at %v\n",
		time.Now().Format(time.UnixDate))
	<-time.After(3 * time.Second)
	fmt.Printf("Code resumed at %v\n",
		time.Now().Format(time.UnixDate))

}
