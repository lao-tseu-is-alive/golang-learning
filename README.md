# Go learning by example

Small, focused programs for learning and experimenting with Go. The repository
targets **Go 1.27**.

## Who this repository is for

The examples are useful after completing the official Go tutorial and are
organized by topic: strings, files, HTTP, concurrency, databases, graphics,
testing, and more. Most files are independent executable examples rather than
packages that form one application.

Start with the standard-library topics (`functions`, `struct`, `strings`,
`numbers`, `time`, `files`, `filesystem`, `testing`, and `concurrency`) before
trying examples that need third-party libraries or external services.

## New to Go? Start with the official resources

If Go is your first compiled language, or if you have never written a Go
program before, begin with the material maintained by the Go project:

1. Visit [Learn Go](https://go.dev/learn/) to install Go and follow the
   official getting-started tutorials.
2. Complete [A Tour of Go](https://go.dev/tour/), the interactive introduction
   to Go's syntax, methods, interfaces, generics, and concurrency. Try the
   exercises rather than only reading the slides.
3. Return here and follow the [learning path](#learning-path) to reinforce each
   concept with small programs you can run, modify, and test locally.

This repository complements the official documentation. It assumes you already
know how to edit a `.go` file and use basic commands such as `go run` and
`go test`.

## Requirements

- [Go 1.27 or newer](https://go.dev/doc/install)
- Git
- Optional system libraries for OpenGL, OpenCV, and raw-network examples
- Optional PostgreSQL instance for database examples

## Get started

Clone the repository and download its modules:

```sh
git clone https://github.com/lao-tseu-is-alive/golang-learning.git
cd golang-learning
go mod download
```

## Learning path

You do not need to complete every example. Follow the stages in order, choose a
few examples from each theme guide, and move forward when you can explain the
milestone in your own words.

| Stage | Focus | Theme guides | Milestone |
| --- | --- | --- | --- |
| 1. Language foundations | Functions, values, methods, structs, interfaces, text, and Unicode | [`functions`](functions/README.md), [`struct`](struct/README.md), [`strings`](strings/README.md), then [`numbers`](numbers/README.md) | Write small functions, define a type with methods, and safely process UTF-8 text. |
| 2. Data and I/O | Readers, writers, files, serialization, paths, and time | [`files`](files/README.md), [`filesystem`](filesystem/README.md), and [`time`](time/README.md), then `input/` and `output/` | Read and write data while checking errors and closing resources. |
| 3. Verification | Unit tests, table-driven tests, examples, and benchmarks | [`testing`](testing/README.md), then `benchmark/` | Test a package and understand a useful failure message. |
| 4. Concurrency | Goroutines, channels, cancellation, and synchronization | [`concurrency`](concurrency/README.md) | Coordinate finite concurrent work without sleeps, leaks, races, or deadlocks. |
| 5. Programs and services | Processes, TCP/UDP, HTTP clients, servers, and WebSockets | [`process`](process/README.md), [`network`](network/README.md), and [`http`](http/README.md) | Build a small service and shut it down cleanly. |
| 6. Integrations | Databases, storage patterns, crypto, graphics, WebAssembly, and containers | `database/`, `pattern_datastore/`, `crypto/`, `opencv/`, `opengl/`, `wasm/`, and `container/` | Recognize external dependencies and keep integration code isolated and testable. |

A short first pass through the supported core is:

```sh
go run ./functions/function-values
go run ./struct/struct-literals
go run ./struct/methods
go run ./strings/reader
go run ./strings/runes-unicode
go run ./files/readfile README.md
go test -v ./testing/...
go run ./concurrency/channels
go run ./concurrency/goroutines-with-waitgroup
go test ./...
```

The theme guides provide the recommended order, prerequisites, and warnings for
examples that intentionally fail, run continuously, or change files.

## Choose a theme

Each top-level directory is a theme. Every runnable example inside it has its
own directory and a conventional `main.go`.

| If you want to learn about... | Start here |
| --- | --- |
| Functions, interfaces, structs, and generics | `functions/`, `struct/`, `generics/`, `geometry/` |
| Text, numbers, dates, and data conversion | `strings/`, `numbers/`, `time/` |
| Input, output, files, and the operating system | `input/`, `output/`, `files/`, `filesystem/`, `process/`, `system/` |
| Goroutines, channels, synchronization, and cancellation | `concurrency/` |
| HTTP clients, servers, routing, and WebSockets | `http/` |
| TCP, UDP, DNS, ICMP, SMTP, and RPC | `network/` |
| SQL, PostgreSQL, and storage abstractions | `database/`, `pattern_datastore/` |
| Hashing, password handling, and encryption | `hash/`, `crypto/` |
| Tests and benchmarks | `testing/`, `benchmark/` |
| Graphics, computer vision, WebAssembly, and containers | `image/`, `opengl/`, `opencv/`, `wasm/`, `container/` |
| Arguments, environment, paths, and runtime information | `getarguments/`, `getenvvar/`, `getcwd/`, `getversion/` |

For example:

```text
concurrency/                 topic
├── channels/               runnable example
│   └── main.go
├── mutex/                  runnable example
│   └── main.go
└── lines-errgroup/         runnable example
    └── main.go
```

## Run examples

Run an individual example by naming its directory. Commands in this README are
intended to be run from the repository root:

```sh
go run ./functions/function-closures
go run ./strings/runes-unicode
go run ./concurrency/channels
```

Multi-file examples work exactly the same way because all of their source files
are kept together:

```sh
go run ./http/websocket-chat
```

Compile every portable example and run all default tests:

```sh
go test ./...
```

Some examples need native libraries and are excluded from the default check by
build tags:

```sh
go run -tags opengl ./opengl/glfw-basic
go run -tags opencv ./opencv/resize-image
go run -tags libpcap ./network/gokping
```

The required native library must be installed before using its tag. The
historical GXUI example has the `gxui` tag but is known not to compile with the
current graphics stack.

## Repository conventions and limitations

Examples under `opencv`, `opengl`, `container`, `database`, and parts of
`network` are advanced integrations. They can require native libraries,
elevated network privileges, Linux, or a running service. Read the source before
running them.

`image/gxui-legacy` is retained for historical reference because GXUI is no
longer maintained. It is not part of the supported example set.

## Maintenance direction

This repository is being modernized as a curated learning collection. New or
renovated examples should:

- keep one directory per runnable example;
- prefer the standard library when it teaches the same concept clearly;
- include a short purpose and run command;
- avoid hidden external prerequisites;
- include deterministic tests when practical; and
- compile under the Go version declared in `go.mod`.

`go test ./...` is the default repository health check. Tagged native examples
should be checked separately on environments that provide their dependencies.
