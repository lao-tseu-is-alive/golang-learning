package main

import (
	"fmt"
	"sync"
)

var names = []string{"Alan", "Joe", "Jack", "Ben",
	"Ellen", "Lisa", "Carl", "Steve",
	"Anton", "Yo"}

type SyncList struct {
	m     sync.Mutex
	slice []interface{}
}

func NewSyncList(cap int) *SyncList {
	return &SyncList{
		sync.Mutex{},
		make([]interface{}, cap),
	}
}

func (l *SyncList) Load(i int) interface{} {
	l.m.Lock()
	defer l.m.Unlock()
	return l.slice[i]
}

func (l *SyncList) Append(val interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	l.slice = append(l.slice, val)
}

func (l *SyncList) Store(i int, val interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	l.slice[i] = val
}

/*
The synchronization primitive Mutex is provided by the package sync.
The Mutex works as a lock above the secured section or resource.
Once the goroutine calls Lock on the Mutex and the Mutex is in the unlocked state,
the Mutex becomes locked and the goroutine gets exclusive access to the critical section.
In case the Mutex is in the locked state, the goroutine calls the Lock method.
This goroutine is blocked and needs to wait until the Mutex gets unlocked again.

Note that in the example, we use the Mutex to synchronize access on a slice primitive,
which is considered to be unsafe for the concurrent use.

The important fact is that the Mutex cannot be copied after its first use.
*/

func main() {

	l := NewSyncList(0)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			l.Append(names[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Printf("Val: %v stored at idx: %d\n", l.Load(i), i)
	}

}
