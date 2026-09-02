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

Run an individual example by naming its file:

```sh
go run ./functions/function_closures.go
go run ./strings/runesUnicode.go
go run ./concurrency/channels.go
```

Some examples consist of several files. Pass all of their files to `go run`:

```sh
go run ./http/websocket_chat_main.go ./http/websocket_chat_hub.go \
  ./http/websocket_chat_client.go
```

Run the packages that currently have maintained tests:

```sh
go test ./testing/math_operation ./benchmark ./pattern_datastore/...
```

## Repository conventions and limitations

Many topic directories intentionally contain multiple `package main` files.
Consequently, `go test ./...` and `go run ./functions` try to combine unrelated
examples and are not supported yet. Run examples by file as shown above.

Examples under `opencv`, `opengl`, `container`, `database`, and parts of
`network` are advanced integrations. They can require native libraries,
elevated network privileges, Linux, or a running service. Read the source before
running them.

`image/test_gxui.go` is a historical example built on the unmaintained GXUI
project. It is retained for reference but does not compile with the current
graphics dependency stack. It should be replaced or moved to a legacy archive.

## Maintenance direction

This repository is being modernized as a curated learning collection. New or
renovated examples should:

- use one directory per runnable example;
- prefer the standard library when it teaches the same concept clearly;
- include a short purpose and run command;
- avoid hidden external prerequisites;
- include deterministic tests when practical; and
- compile under the Go version declared in `go.mod`.

The long-term goal is to make `go test ./...` a reliable health check after the
single-file examples have been separated into their own directories.
