package main

import "sync"
import "fmt"

/*
While working with concurrently running code branches, it is no exception that at some point
the program needs to wait for concurrently running parts of the code.
This recipe gives insight into how to use the WaitGroup to wait for running goroutines.

With help of the  WaitGroup struct from the sync package, the program run is able to wait until
some finite number of goroutines finish.
The WaitGroup struct implements the method Add to add the number of goroutines to wait for.
Then after the goroutine finishes,  the Done method should be called to decrement
the number of goroutines to wait for.
The method Wait is called as a block until the given number of Done calls has been done
(usually at the end of a goroutine). The WaitGroup should be used the same way
as all synchronization primitives within the sync package.
After the creation of the object, the struct should not be copied.
*/

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			// Do some work
			defer wg.Done()
			fmt.Printf("Exiting %d\n", idx)
		}(i)
	}
	wg.Wait()
	fmt.Println("All done.")
}
