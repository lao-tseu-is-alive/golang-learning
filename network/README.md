# Networking

This theme covers addresses, DNS, interfaces, TCP, UDP, ICMP, RPC, and a few
self-contained HTTP client/server experiments. Network results depend on the
host, firewall, DNS, and remote services. Only probe systems you own or are
explicitly authorized to test.

[Back to the learning path](../README.md#learning-path)

## Mental model

The `net` package provides common abstractions across network transports. An
address identifies an endpoint. A `net.Conn` is a bidirectional connection; a
`net.Listener` accepts new connections. Always close connections and listeners,
and use deadlines or contexts so unavailable peers cannot block forever.

TCP is a reliable ordered byte stream. It does not preserve message boundaries,
so applications need framing such as delimiters, fixed lengths, or a structured
protocol. UDP sends independent datagrams without delivery or ordering
guarantees. DNS maps names and addresses through external infrastructure and
can return multiple results or fail temporarily.

HTTP and RPC are application protocols built above transports such as TCP.
Prefer their higher-level packages unless the lesson specifically concerns the
wire protocol. Tests should usually use loopback addresses, ephemeral ports,
and local test servers instead of public Internet services.

ICMP and packet capture may require operating-system configuration, native
libraries, or elevated privileges. Port scanning can be disruptive and must be
limited to authorized targets.

## Core path

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`url`](url/) | Building, parsing, and serializing a URL | Local computation only |
| [`headers`](headers/) | HTTP header values, replacement, and deletion | Local computation only |
| [`get-net-interfaces`](get-net-interfaces/) | Enumerating interfaces and assigned addresses | Output depends on the host |
| [`tcpclient`](tcpclient/) | Speaking HTTP over a raw TCP connection | Self-contained; uses localhost port 7070 |
| [`http-request`](http-request/) | Building a request and serving it locally | Self-contained; uses localhost port 7070 |
| [`restful`](restful/) | JSON request/response flow over local HTTP | Self-contained; uses localhost port 7070 |
| [`jsonrpc`](jsonrpc/) | A local JSON-RPC server and client | Uses `golog` and localhost port 8222 |
| [`json-rpc`](json-rpc/) | Several local RPC methods and errors | Uses localhost port 8222 |

```sh
go run ./network/url
go run ./network/headers
go run ./network/get-net-interfaces
go run ./network/tcpclient
go run ./network/restful
go run ./network/json-rpc
```

## Servers and external-network examples

| Example | What it teaches | Requirements or behavior |
| --- | --- | --- |
| [`server-echo`](server-echo/) | A sequential TCP echo server | Runs on port 8080; use `nc localhost 8080` |
| [`server-tcp-multiple-clients`](server-tcp-multiple-clients/) | A goroutine per TCP connection | Runs continuously on port 8080 |
| [`server-echo-udp`](server-echo-udp/) | UDP datagrams and peer addresses | Runs on port 8080; send `STOP` to exit |
| [`http-basic-time-server`](http-basic-time-server/) | HTTP server timeouts and a JSON response | Runs continuously on port 8080 |
| [`httpdemo`](httpdemo/) | Two forms of local HTTP POST | Similar to `http-request`; uses older `ioutil` APIs |
| [`dnslookup`](dnslookup/) | Forward and reverse DNS lookup | Uses fixed local and public addresses |
| [`check-websites-status-serial-execution`](check-websites-status-serial-execution/) | Sequential HTTP checks and timing | Contacts fixed public sites; no client timeout |
| [`check-websites-status-concurrent-execution`](check-websites-status-concurrent-execution/) | Concurrent HTTP checks and aggregation | Use `-urlFile network/urls.txt`; contacts public sites |
| [`redirects`](redirects/) | Client redirect policy and request history | Self-contained localhost server |
| [`port-scanner`](port-scanner/) | Bounded concurrent TCP connection attempts | Scans 65,535 ports; use only on an authorized address |

## Privileged or historical integrations

| Example | What it teaches | Requirements or caution |
| --- | --- | --- |
| [`fastping`](fastping/) | ICMP callbacks with `go-fastping` | Usually requires raw-socket privileges |
| [`goping01`](goping01/) | A finite ping using `go-ping` | Linux may require `ping_group_range` configuration |
| [`goping02`](goping02/) | Ping callbacks and Ctrl-C handling | Runs until stopped; OS privileges vary |
| [`gokping`](gokping/) | Packet capture and multi-host ping statistics | Requires `-tags libpcap`, libpcap, and authorized targets |
| [`smtp`](smtp/) | Manual SMTP plus STARTTLS | Historical; hard-coded provider/config assumptions; do not use real credentials here |
| [`smtp-simple-read-ini`](smtp-simple-read-ini/) | SMTP helper plus INI configuration | Historical; hard-coded provider/config assumptions |

## Check your understanding

- Explain why one TCP `Read` is not guaranteed to return one application
  message.
- Add a timeout and always close a response body in an HTTP client.
- Make a local server accept an ephemeral port so tests can run concurrently.
- Compare the delivery guarantees of TCP and UDP.

Continue with [`http`](../http/README.md).

## Further reading

- [`net` package](https://pkg.go.dev/net)
- [`net/url` package](https://pkg.go.dev/net/url)
- [`net/http` package](https://pkg.go.dev/net/http)
- [Go concurrency patterns: context](https://go.dev/blog/context)
