package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Getting the fastest result from multiple sources
In some cases, for example, while integrating information retrieval from multiple sources,
you only need the first result, the fastest one, and the other results are irrelevant after that.
An example from the real world could be extracting the currency rate to count the price.
You have multiple third-party services and because you need to show the prices as fast as possible,
you need only the first rate received from any service.
This recipe will show the pattern for how to achieve such behavior.

The preceding code proposes the solution on executing multiple tasks that output some results,
and we need only the first fastest one.
The solution uses the Context with the cancel function to call cancel once
the first result is obtained.
The SearchSrc structure provides the Search method that results in a  channel where the result is written.
Note that the Search method simulates the delay with the time.Sleep function.
The merge function, for each channel from the Search method,
triggers the goroutine that writes to the final output channel that is read in the main method.
While the first result is received from the output channel produced from the merge function,
the CancelFunc stored in the variable cancel is called to cancel the rest of the processing.
*/

type SearchSrc struct {
	ID    string
	Delay int
}

func (s *SearchSrc) Search(ctx context.Context) <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(time.Duration(s.Delay) * time.Second)
		select {
		case out <- "Result " + s.ID:
		case <-ctx.Done():
			fmt.Println("Search received Done()")
		}
		close(out)
		fmt.Println("Search finished for ID: " + s.ID)
	}()
	return out
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	src1 := &SearchSrc{"1", 2}
	src2 := &SearchSrc{"2", 6}

	r1 := src1.Search(ctx)
	r2 := src2.Search(ctx)

	out := merge(ctx, r1, r2)

	for firstResult := range out {
		cancel()
		fmt.Println("First result is: " + firstResult)
	}
}

func merge(ctx context.Context, results ...<-chan string) <-chan string {
	wg := sync.WaitGroup{}
	out := make(chan string)

	output := func(c <-chan string) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Received ctx.Done()")
		case res := <-c:
			out <- res
		}
	}

	wg.Add(len(results))
	for _, c := range results {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
