# Functions

This theme introduces functions as first-class values and then shows how a
closure retains state between calls. It is the best starting point after basic
Go syntax. Run commands from the repository root.

[Back to the learning path](../README.md#learning-path)

## Mental model

A function gives a name and a contract to a piece of behavior. Its signature
describes what it needs and what it returns; callers should not need to know how
the result is produced. Small functions are easier to read, combine, and test.

In Go, functions are also values. A function can be assigned to a variable,
passed as an argument, or returned from another function. This lets code accept
behavior without introducing a new type for every variation.

A closure is a function value that refers to variables outside its own body.
Those captured variables remain available after the surrounding function has
returned. This is useful for configuration and small stateful behaviors, but
remember that the state is shared by every call to that particular closure. If
multiple goroutines call it, that state may need synchronization.

As you read the examples, ask three questions:

- What is the function's input/output contract?
- Is behavior being passed around as data?
- Does the function retain or mutate hidden state?

## Recommended order

| Example | What it teaches | Level |
| --- | --- | --- |
| [`function-values`](function-values/) | Passing functions as arguments and using anonymous functions | Beginner |
| [`function-closures`](function-closures/) | Capturing and updating variables from an enclosing scope | Beginner |
| [`fibonacci-closure`](fibonacci-closure/) | Applying a stateful closure to generate a sequence | Beginner |

```sh
go run ./functions/function-values
go run ./functions/function-closures
go run ./functions/fibonacci-closure
```

## Check your understanding

- Change `compute` so it accepts the two input values instead of fixing them.
- Create two independent counters with the same closure factory and explain why
  they do not share state.
- Add a test for the first ten values returned by the Fibonacci closure.

Continue with [`struct`](../struct/README.md).

## Further reading

- [Functions](https://go.dev/tour/basics/4) in A Tour of Go
- [Function values](https://go.dev/tour/moretypes/24)
- [Function closures](https://go.dev/tour/moretypes/25)
