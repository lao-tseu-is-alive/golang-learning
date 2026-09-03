# Processes and signals

This theme starts external programs, connects their streams, observes process
state, and responds to operating-system signals. Several examples assume Unix
commands or signals and cannot run in the Go Playground.

[Back to the learning path](../README.md#learning-path)

## Mental model

A goroutine is concurrent work inside the current Go process. A child process
is a separate operating-system program with its own address space, environment,
standard streams, exit status, and lifecycle.

`exec.Command` executes a program directly; it does not invoke a shell. Shell
features such as pipes, redirects, wildcard expansion, and quoting are therefore
not interpreted. Pass arguments as separate strings. If an application chooses
to invoke a shell, untrusted input must never be concatenated into a command.

`Run` starts a command and waits. `Start` returns after starting it, and must be
paired with `Wait` to collect its exit status and release resources. `Output`,
`CombinedOutput`, and the stdin/stdout/stderr pipe methods cover common data-flow
patterns. Use `CommandContext` when a deadline or cancellation should terminate
work that is no longer useful.

Signals are platform-specific notifications. They can request graceful
shutdown, but some signals—such as Unix `SIGKILL`—cannot be caught. Cleanup must
therefore be useful but not the only mechanism protecting important data.

## Recommended order

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`echo-stdin`](echo-stdin/) | Reading a child or parent process's standard input | Pipe text into it |
| [`run-command-sync`](run-command-sync/) | Starting a command and waiting with `Run` | Uses `ls`; primarily Unix-oriented |
| [`run-command-async`](run-command-async/) | Separating `Start` from `Wait` | Uses `ls`; primarily Unix-oriented |
| [`send-strings-to-another-process`](send-strings-to-another-process/) | Connecting parent and child stdin/stdout pipes | Starts the local `echo-stdin` example, then terminates it |
| [`child-process-id`](child-process-id/) | PID, process state, `Wait`, context, and call frames | Requires `sleep` on Unix or `timeout` on Windows; uses `golog` |

```sh
printf 'hello from stdin\n' | go run ./process/echo-stdin
go run ./process/run-command-sync
go run ./process/run-command-async
go run ./process/send-strings-to-another-process
```

## Environment-specific examples

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`getprocessid`](getprocessid/) | Current PID and external process inspection | Requires the Unix `ps` command |
| [`signals`](signals/) | Receiving selected Unix signals | Runs until signalled; signal set is Unix-oriented |
| [`signal-handling`](signal-handling/) | Coordinating shutdown and resource cleanup | Runs until signalled and creates a timestamped log file |
| [`get-stdout-from-process`](get-stdout-from-process/) | Scanning a child process's stdout with a timeout | Requires `ping` and network/DNS access; output format is platform-dependent |

## Check your understanding

- Explain why `Start` without `Wait` leaks process resources.
- Capture stderr separately and report a child's exit status.
- Cancel a long-running child with a context and distinguish timeout from a
  normal non-zero exit.
- Design shutdown so correctness does not depend on receiving a catchable
  signal.

Continue with [`network`](../network/README.md).

## Further reading

- [`os/exec` package](https://pkg.go.dev/os/exec)
- [`os.Process` documentation](https://pkg.go.dev/os#Process)
- [`os/signal` package](https://pkg.go.dev/os/signal)
- [Go command and path security](https://go.dev/blog/path-security)
