package main

import "sync"
import "fmt"
import "time"

type Worker struct {
	id string
}

func (w *Worker) String() string {
	return w.id
}

var globalCounter = 0

var pool = sync.Pool{
	New: func() interface{} {
		res := &Worker{fmt.Sprintf("%d", globalCounter)}
		globalCounter++
		return res
	},
}

/*
The sync package contains the struct for pooling the resources.
The Pool struct has the Get and Put method to retrieve and put the resource back to the pool.
The Pool struct is considered to be safe for concurrent access.

While creating the Pool struct, the New field needs to be set.
The New field is a no-argument function that should return the pointer to the pooled item.
This function is then called in case the new object in the pool needs to be initialized.

Note from the logs of the preceding example, that the Worker is reused while returned to the pool.
The important fact is that there shouldn't be any assumption related to the retrieved items
by Get and returned items to Put method (like I've put three objects to pool just now,
so there will be at least three available).
This is mainly caused by the fact that that the idle items in a Pool could be automatically removed at any time.
*/

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			// This code block is done only once
			w := pool.Get().(*Worker)
			fmt.Println("Got worker ID: " + w.String())
			time.Sleep(time.Second)
			pool.Put(w)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
