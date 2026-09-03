# Testing

This theme demonstrates Go's built-in test runner, helpers, subtests,
table-driven cases, executable examples, and `TestMain`. Unlike most themes,
these examples are run with `go test`, not `go run`.

[Back to the learning path](../README.md#learning-path)

## Mental model

A test is another Go function that states an observable promise about the code.
Good tests focus on behavior: given this input and state, expect this result or
error. They should be deterministic, isolated, and clear enough that a failure
helps locate the broken promise.

Go deliberately keeps the basic testing model small. Files end in `_test.go`,
test functions receive `*testing.T`, and `go test` builds a temporary test
binary. Table-driven tests express many input/output cases as data; `t.Run`
gives each case a useful name and its own test context.

Use `t.Error` or `t.Errorf` when the current test can safely continue. Use
`t.Fatal` or `t.Fatalf` when continuing would be meaningless or unsafe. Mark
shared assertion helpers with `t.Helper` so failures point to the caller.

Executable examples verify documentation output. Benchmarks answer performance
questions, but only after correctness is established. The race detector checks
for unsynchronized memory access during the paths that actually execute, so it
complements rather than replaces tests.

## Recommended order

| Example | What it teaches | Command |
| --- | --- | --- |
| [`sample_test.go`](sample_test.go) | Basic tests, fatal failures, helpers, logging, and skips | `go test -v ./testing` |
| [`subtest_test.go`](subtest_test.go) | Named subtests and looping over test data | `go test -v ./testing -run TestSampleSubtest` |
| [`math_operation`](math_operation/) | Package tests, table-driven cases, examples, and `TestMain` | `go test -v ./testing/math_operation` |
| [`benchmark`](../benchmark/) | Benchmark functions and the `-bench` flag | `go test -bench=. ./benchmark` |

The first two files contain deliberately failing demonstrations guarded by
`t.Skip`. Remove a guard locally when you specifically want to compare failure
locations or observe failing subtests.

## Useful commands

```sh
go test -v ./testing/...
go test -cover ./testing/...
go test -race ./testing/...
go test -bench=. ./benchmark
```

## Check your understanding

- Convert a repeated set of assertions into a table-driven test using `t.Run`.
- Mark a shared assertion function with `t.Helper` and compare the failure
  location before and after.
- Add one executable `Example` with an `// Output:` comment.
- Explain what the race detector checks that a normal unit test does not.

Continue with [`concurrency`](../concurrency/README.md).

## Further reading

- [Add a test](https://go.dev/doc/tutorial/add-a-test), the official tutorial
- [`testing` package](https://pkg.go.dev/testing)
- [Table-driven tests using subtests](https://go.dev/blog/subtests)
- [Data race detector](https://go.dev/doc/articles/race_detector)
