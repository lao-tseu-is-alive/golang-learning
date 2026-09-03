# Strings, text, and Unicode

Go strings contain bytes and UTF-8 text is processed as runes. Start with the
standard-library examples, then explore Unicode normalization, encodings,
templates, and command-line text filters. Run commands from the repository
root so bundled input paths resolve correctly.

[Back to the learning path](../README.md#learning-path)

## Mental model

A Go string is an immutable sequence of bytes. It often contains UTF-8 text,
but the type itself does not guarantee valid UTF-8. Consequently, `len(s)`
counts bytes and `s[i]` returns one byte—not necessarily one visible character.

A `rune` is an alias for `int32` used to represent a Unicode code point. A
`for range` loop over a string decodes one UTF-8 rune at a time and reports its
starting byte position. Even a rune is not always what a user perceives as one
character: an accented letter or emoji can contain several code points.

Use the representation that matches the job:

- keep `string` for immutable text and arbitrary byte content;
- use `[]byte` for byte-oriented I/O, protocols, or mutable binary data;
- use `[]rune` only when code-point indexing or mutation is genuinely needed;
- use `strings.Reader` or another `io.Reader` when text should flow through a
  streaming API.

Text comparison can also require case folding or Unicode normalization. Do not
assume that lowercasing or comparing bytes solves every human-language problem.

## Core path

| Example | What it teaches | Level |
| --- | --- | --- |
| [`reader`](reader/) | Reading a string through the `io.Reader` interface | Beginner |
| [`stringreplacer`](stringreplacer/) | Reusable multi-string replacement | Beginner |
| [`text-whitespace-align`](text-whitespace-align/) | Trimming, collapsing, and padding text | Beginner |
| [`runes-unicode`](runes-unicode/) | Bytes versus runes, Unicode properties, and normalization | Intermediate; uses `golang.org/x/text` |
| [`read-csv`](read-csv/) | Reading structured records with `encoding/csv` | Intermediate; uses bundled `strings/data.csv` |

```sh
go run ./strings/reader
go run ./strings/stringreplacer
go run ./strings/text-whitespace-align
go run ./strings/runes-unicode
go run ./strings/read-csv
```

## More text-processing examples

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`string-case`](string-case/) | Case conversion and case-insensitive comparison | Contains deprecated `strings.Title` usage; modernization candidate |
| [`getwords`](getwords/) | A Unicode-aware Unix-style stdin filter | Uses `golog`; pipe UTF-8 text into stdin |
| [`splitwords`](splitwords/) | Scanning lines and extracting words with regular expressions | Requires a text file supplied with `-file` |
| [`wordcount-test`](wordcount-test/) | A self-contained word-count exercise and test harness | Educational predecessor to normal `go test` |
| [`string-regex-replace`](string-regex-replace/) | Capture groups and regular-expression replacement | Uses `golog` |
| [`text-ident`](text-ident/) | Adding and removing indentation | Standard library |
| [`tabwriter`](tabwriter/) | Producing aligned tabular text | Standard library |
| [`string-template`](string-template/) | `text/template`, blocks, cloning, and functions | Standard library |
| [`html-template`](html-template/) | Context-aware HTML templates | Standard library |
| [`go-read-win1252-file`](go-read-win1252-file/) | Decoding Windows-1252 text into UTF-8 | Uses `golang.org/x/text`, `goutils`, and bundled input |
| [`write-win1252-file`](write-win1252-file/) | Encoding text with a legacy Windows character set | Creates `Windows1250.txt` in the working directory |

Example filter invocation:

```sh
printf 'Go makes Unicode text pleasant.\n' | go run ./strings/getwords -min-length=3
go run ./strings/splitwords -file strings/win1252.txt
```

## Check your understanding

- Explain why `len(s)` and `utf8.RuneCountInString(s)` can differ.
- Change the CSV example to report malformed records without losing the valid
  records already read.
- Turn one text transformation into a package function with table-driven tests.

Continue with [`files`](../files/README.md).

## Further reading

- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [`strings` package](https://pkg.go.dev/strings)
- [`unicode/utf8` package](https://pkg.go.dev/unicode/utf8)
- [Text normalization in Go](https://go.dev/blog/normalization)
