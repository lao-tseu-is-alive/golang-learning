# Files and serialization

This theme covers common `io` patterns, filesystem access, and standard data
formats. Run commands from the repository root. Some examples deliberately
create or modify files; those side effects are called out below.

[Back to the learning path](../README.md#learning-path)

## Mental model

File handling becomes simpler when separated into two concerns: where bytes
come from or go to, and how those bytes are interpreted. Go expresses the first
concern mainly through the tiny `io.Reader` and `io.Writer` interfaces. The same
copying or decoding code can therefore work with a file, memory buffer, network
connection, compressed stream, or HTTP body.

Choose whole-file helpers for small, bounded inputs. Use streaming readers and
writers when inputs can be large, arrive gradually, or should be transformed
without keeping everything in memory.

Files are operating-system resources and operations can fail at every stage.
Open them with an explicit purpose, check errors, close them when finished, and
avoid overwriting user data by surprise. Also distinguish a path from its
contents: relative paths are resolved from the program's current working
directory, not from the directory containing `main.go`.

Serialization packages such as `encoding/json`, `encoding/xml`, and
`encoding/gob` translate between byte streams and Go values. Validate decoded
data separately; successful decoding does not mean the data is meaningful for
your application.

## Core path

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`readfile`](readfile/) | Opening and reading a complete file | Pass a filename; read-only |
| [`multiwriter`](multiwriter/) | Sending the same bytes to several `io.Writer` values | Modifies `files/sample.txt` |
| [`rwbinary`](rwbinary/) | Encoding and decoding fixed binary values with byte order | In-memory only |
| [`gob`](gob/) | Encoding and decoding Go values with `encoding/gob` | In-memory only |
| [`readxml`](readxml/) | Streaming XML tokens into structs | Reads bundled `files/data.xml` |
| [`zip`](zip/) | Creating and reading a ZIP archive | Creates `data.zip` in the working directory |

```sh
go run ./files/readfile README.md
go run ./files/rwbinary
go run ./files/gob
go run ./files/readxml
```

## Further examples

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`allfiles`](allfiles/) | Walking, sorting, and listing a directory tree | Use `-dir` to limit the scanned tree |
| [`fileseek`](fileseek/) | Random access with `Seek`, `ReadAt`, and `WriteAt` | Creates or modifies `flatfile.txt` |
| [`writetestfile`](writetestfile/) | Opening and writing a file | Creates or modifies `test.txt` |
| [`charset`](charset/) | Encoding and decoding Windows-1252 text | Uses `golang.org/x/text`; creates `example.txt` |
| [`pipe`](pipe/) | Connecting an external process through `io.Pipe` | Requires an `echo` executable |
| [`readprogressive`](readprogressive/) | Comparing scanner-based and whole-file reading | Requires an external sample text under `data/` |
| [`json`](json/) | Token-based JSON decoding and decoder error behavior | The embedded JSON is intentionally incomplete |

Examples that write files are best run from a temporary working directory if
you do not want to change the checkout. The run commands above are read-only or
in-memory.

## Check your understanding

- Replace a whole-file read with streaming when the input could be large.
- Make one writing example accept an output path and avoid overwriting an
  existing file by default.
- Add round-trip tests for one serialization format.

Continue with [`testing`](../testing/README.md).

## Further reading

- [`io` package and its core interfaces](https://pkg.go.dev/io)
- [`os` package](https://pkg.go.dev/os)
- [Effective Go: interfaces](https://go.dev/doc/effective_go#interfaces_and_types)
- [JSON and Go](https://go.dev/blog/json)
