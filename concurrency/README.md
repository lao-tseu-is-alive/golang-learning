# Concurrency

This theme progresses from goroutines to channels, cancellation, and the
`sync` package. Concurrent output order may change between runs. Begin with the
finite examples below; two demonstrations intentionally do not terminate
normally.

[Back to the learning path](../README.md#learning-path)

## Mental model

Concurrency and parallelism are related, but they answer different questions:

- **Concurrency** is how a program is structured to make progress on multiple
  independent activities during the same period of time.
- **Parallelism** is when computations literally execute at the same instant,
  usually on different CPU cores.

A concurrent program can run on one core by interleaving work. A parallel
program needs hardware and a runtime able to execute work simultaneously.
Correct concurrent code must not depend on whether the runtime happens to run
its goroutines in parallel.

### Goroutines

A goroutine is a function executing concurrently with other goroutines in the
same process. Start one with the `go` keyword. Goroutines are managed by the Go
runtime and multiplexed over operating-system threads; they are not promises,
futures, or one-thread-per-task.

Starting a goroutine also creates a lifecycle obligation. Decide how it will
finish, how errors are reported, and how it will be cancelled if its result is
no longer needed. Returning from `main` terminates the process; it does not wait
for other goroutines. Use synchronization instead of sleeps to know when work
is complete.

### Channels and `select`

A channel is a typed communication and synchronization mechanism between
goroutines. Sending and receiving describe when values may move between parts
of the program:

```go
results := make(chan int)
go func() {
	results <- calculate()
}()
result := <-results
```

With an unbuffered channel, a send and receive rendezvous: each waits for the
other. A buffered channel can hold a limited number of values, allowing sender
and receiver to progress at different moments until the buffer is full or
empty. A buffer changes coordination; it is not a general cure for deadlocks.

Closing a channel communicates that no more values will be sent. The sending
side normally owns that decision. Receivers can use `range` or the comma-ok form
to detect completion. Channels do not need to be closed merely to release a
resource.

`select` waits until one of several channel operations can proceed. It is the
foundation for timeouts, cancellation, and coordinating several sources.

### Shared state and cancellation

Channels are excellent when goroutines exchange ownership or stream values.
When several goroutines must protect the same in-memory state, a mutex can be
clearer. Choose the mechanism that makes ownership and invariants easiest to
explain.

Use `context.Context` to propagate deadlines and cancellation across a call
tree. The goroutine that starts work should normally retain the cancel function,
and every started goroutine should have a path that observes cancellation and
returns. Use `go test -race` to detect many—but not all—incorrect shared-memory
accesses.

## Core path

| Example | What it teaches | Level |
| --- | --- | --- |
| [`goroutines`](goroutines/) | Starting a goroutine and observing concurrent execution | Beginner |
| [`goroutines-with-waitgroup`](goroutines-with-waitgroup/) | Waiting for a known number of goroutines | Beginner |
| [`channels`](channels/) | Sending results through an unbuffered channel | Beginner |
| [`basic-channel`](basic-channel/) | Ranging until a sender closes a channel | Beginner |
| [`buffered-channels`](buffered-channels/) | Channel capacity and blocking | Beginner; intentionally ends in deadlock |
| [`select`](select/) | Coordinating send and receive cases | Intermediate |
| [`default-selection`](default-selection/) | Non-blocking `select`, timers, and timeouts | Intermediate |
| [`mutex`](mutex/) | Protecting shared state with `sync.Mutex` | Intermediate |
| [`lines-errgroup`](lines-errgroup/) | Collecting errors from concurrent work | Intermediate; uses `golang.org/x/sync` |
| [`first`](first/) | Taking the fastest result and cancelling remaining work | Intermediate |

```sh
go run ./concurrency/goroutines
go run ./concurrency/goroutines-with-waitgroup
go run ./concurrency/channels
go run ./concurrency/basic-channel
go run ./concurrency/select
go run ./concurrency/default-selection
go run ./concurrency/mutex
go run ./concurrency/lines-errgroup
go run ./concurrency/first
```

## Additional synchronization examples

| Example | What it teaches | Behavior or caution |
| --- | --- | --- |
| [`channels-range-close`](channels-range-close/) | Closing a channel to finish a range loop | Produces 99 Fibonacci values |
| [`basic-select`](basic-select/) | Receiving from independently paced channels | Runs continuously; stop it with Ctrl-C |
| [`mutex-counter`](mutex-counter/) | Protecting a map with a mutex | Uses a sleep to let workers finish; prefer a `WaitGroup` in real code |
| [`map-sync`](map-sync/) | Concurrent access with `sync.Map` | Prefer a map plus mutex unless `sync.Map` fits its documented use case |
| [`once`](once/) | One-time initialization with `sync.Once` | Intermediate |
| [`pool`](pool/) | Reusing temporary objects with `sync.Pool` | Advanced; pool contents may disappear at any time |
| [`syncgroup`](syncgroup/) | Another finite `WaitGroup` pattern | Similar to `goroutines-with-waitgroup` |

Run the intentional deadlock demonstration only when you want to inspect the
runtime diagnostic:

```sh
go run ./concurrency/buffered-channels
```

## Check your understanding

- Replace a sleep used for synchronization with a channel or `sync.WaitGroup`.
- State which goroutine owns a channel and therefore should close it.
- Add cancellation to a worker and verify that it exits when the context is
  done.
- Run `go test -race ./...` after changing code that shares memory.

Continue with `process/`, `network/`, and `http/` from the
[learning path](../README.md#learning-path).

## Further reading and watching

- [Concurrency is not parallelism](https://go.dev/blog/waza-talk), Rob Pike's
  core mental model, with slides and background
- [Concurrency is not parallelism—video](https://www.youtube.com/watch?v=oV9rvDllKEg)
- [Concurrency in A Tour of Go](https://go.dev/tour/concurrency/1)
- [Go concurrency patterns: pipelines and cancellation](https://go.dev/blog/pipelines)
- [Go concurrency patterns: context](https://go.dev/blog/context)
- [The Go memory model](https://go.dev/ref/mem)
- [Data race detector](https://go.dev/doc/articles/race_detector)
