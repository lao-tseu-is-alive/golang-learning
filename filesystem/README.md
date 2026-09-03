# Filesystem

This theme deals with paths, directories, metadata, permissions, temporary
files, traversal, and platform-specific filesystem information. Run commands
from the repository root and read the side-effect column before executing an
example.

[Back to the learning path](../README.md#learning-path)

## Mental model

A path is a name used to locate an entry; it is not the file itself.
`path/filepath` follows operating-system path conventions, while `path` is for
slash-separated paths such as URLs. Relative paths are interpreted from the
process's current working directory.

The `os` package works with the host filesystem. The `io/fs` interfaces let
code operate on other filesystem implementations too, including embedded and
in-memory trees. Prefer the abstraction when the code only needs to read or
walk a filesystem.

Filesystem operations race with the outside world: an entry can change between
`Stat` and `Open`. Treat errors from the operation itself as authoritative.
Choose permissions deliberately, avoid overwriting data unexpectedly, and use
temporary files or directories when intermediate data should be isolated.

Paths, permissions, user information, syscalls, and symbolic links differ
between operating systems. Portable code should avoid assuming Unix semantics
unless that limitation is explicit.

## Core path

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`listdir`](listdir/) | Comparing directory reads with tree walking | Uses older `ioutil.ReadDir`; pass `-dir filesystem` |
| [`filterglob`](filterglob/) | Matching paths with `filepath.Glob` | Creates and removes six temporary example files |
| [`fileinfo-md5`](fileinfo-md5/) | File metadata and streaming a checksum | Uses `golog`; pass a filename |
| [`tempfile`](tempfile/) | Temporary files, temporary directories, and cleanup | Uses older `ioutil` helpers |
| [`homedir`](homedir/) | Looking up the current user and groups | Results depend on the host environment |

```sh
go run ./filesystem/listdir -dir filesystem
go run ./filesystem/filterglob
go run ./filesystem/fileinfo-md5 README.md
go run ./filesystem/tempfile
```

## Mutating and platform-specific examples

| Example | What it teaches | Side effect or limitation |
| --- | --- | --- |
| [`createfilesanddir`](createfilesanddir/) | Creating files and nested directories | Writes several entries and may fail when rerun |
| [`writefile`](writefile/) | Writing strings and copying from a reader | Creates or truncates `sample.file` |
| [`filechmod`](filechmod/) | Inspecting and changing permission bits | Creates `test.txt`; Unix permissions differ from Windows |
| [`filecompare`](filecompare/) | Checksum and line-by-line comparison | Creates and then removes three files; MD5 is not collision-resistant |
| [`syncwrite`](syncwrite/) | Serializing concurrent writes with a mutex | Creates `sample_sync_write.txt` |
| [`df-linux-sys-unix`](df-linux-sys-unix/) | Disk and inode usage with `x/sys/unix` | Non-Windows; recommended syscall variant |
| [`df-linux-syscall`](df-linux-syscall/) | Historical direct `syscall.Statfs` usage | Non-Windows; prefer `x/sys/unix` |

## Check your understanding

- Predict which directory a relative output path will modify.
- Rewrite an operation to accept an `fs.FS` so it can be tested in memory.
- Make a mutating example idempotent and ensure cleanup still happens on error.
- Explain why checking for existence before opening does not eliminate races.

Continue with [`process`](../process/README.md).

## Further reading

- [`os` package](https://pkg.go.dev/os)
- [`path/filepath` package](https://pkg.go.dev/path/filepath)
- [`io/fs` filesystem interfaces](https://pkg.go.dev/io/fs)
- [Go path security](https://go.dev/blog/path-security)
