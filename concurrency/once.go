package main

import (
	"fmt"
	"sync"
)

var names = []interface{}{"Alan", "Joe", "Jack", "Ben",
	"Ellen", "Lisa", "Carl", "Steve",
	"Anton", "Yo"}

/*
The sample code illustrates the lazy loading of the data while accessing the container structure.
As the data should be loaded only once, the Once struct from the sync package is used in the method Pop.
The Once implements only one method called Do which consumes the func with no arguments
and the function is executed only once per Once instance, during the execution.

The Do method calls blocks until the first run is done.
This fact corresponds with the fact that Once is intended to be used for initialization.
*/
type Source struct {
	m    *sync.Mutex
	o    *sync.Once
	data []interface{}
}

func (s *Source) Pop() (interface{}, error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.o.Do(func() {
		s.data = names
		fmt.Println("Data has been loaded.")
	})
	if len(s.data) > 0 {
		res := s.data[0]
		s.data = s.data[1:]
		return res, nil
	}
	return nil, fmt.Errorf("No data available")
}

func main() {

	s := &Source{&sync.Mutex{}, &sync.Once{}, nil}
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			// This code block is done only once
			if val, err := s.Pop(); err == nil {
				fmt.Printf("Pop %d returned: %s\n", idx, val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
