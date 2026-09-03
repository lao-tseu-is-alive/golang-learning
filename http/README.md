# HTTP clients and servers

This theme progresses from a small `net/http` server to routing, JSON, forms,
middleware, graceful shutdown, files, sessions, TLS, HTTP/2, and WebSockets.
Most servers listen on port 8080 and run until interrupted; start only one such
example at a time.

[Back to the learning path](../README.md#learning-path)

## Mental model

HTTP is a request/response protocol. A client sends a method, URL, headers, and
possibly a body; the server returns a status, headers, and possibly a body. In
Go, `http.Request` represents the request and `http.ResponseWriter` builds the
response.

The central server abstraction is `http.Handler`: anything with
`ServeHTTP(ResponseWriter, *Request)` can handle a request. `HandlerFunc` adapts
an ordinary function to that interface, and `ServeMux` selects a handler from a
route pattern. Middleware is simply a handler that performs work before or
after delegating to another handler.

Status and headers must be chosen before writing the body. Treat every path,
query, header, form field, cookie, and uploaded byte as untrusted input. Apply
size limits, validate data, escape output in the correct context, and never use
demonstration keys or authentication checks in production.

Clients should normally have a timeout, close response bodies, and inspect
status codes. Servers should set appropriate timeouts, propagate
`r.Context()`, report errors without panicking inside handlers, and shut down
gracefully so in-flight requests can finish.

## Core path

| Example | What it teaches | How to explore it |
| --- | --- | --- |
| [`simple-server`](simple-server/) | Handlers and static resources | Open `http://localhost:8080` |
| [`server`](server/) | Inspecting method, URL, headers, host, and peer | Open or curl localhost; uses `golog` |
| [`http-mux-router`](http-mux-router/) | Standard-library sub-routing and prefixes | Try `/user`, `/time`, `/items/clothes`, and `/admin/ports` |
| [`http-json`](http-json/) | Streaming JSON decoding and encoding | Use the curl commands documented in `main.go` |
| [`html-template`](html-template/) | Rendering embedded HTML safely | Open localhost |
| [`html-form`](html-form/) | GET/POST form handling with an embedded template | Submit the form in a browser |
| [`http-post`](http-post/) | Parsing URL and form parameters | Use the printed curl command |
| [`http-middleware`](http-middleware/) | Wrapping handlers and composing middleware | Send GET and non-GET requests |
| [`graceful-shutdown`](graceful-shutdown/) | Stopping a server with context and Ctrl-C | Interrupt after making a request |

```sh
go run ./http/simple-server
go run ./http/http-mux-router
go run ./http/http-json
go run ./http/html-template
go run ./http/http-middleware
go run ./http/graceful-shutdown
```

Run only one command at a time and stop it with Ctrl-C before starting the next.

## Common web features

| Example | What it teaches | Requirements or caution |
| --- | --- | --- |
| [`gocurl`](gocurl/) | A minimal HTTP GET client | Contacts an external service; uses older `ioutil.ReadAll` |
| [`http-basic-file-serve`](http-basic-file-serve/) | `FileServer`, `StripPrefix`, and `ServeFile` | Serves bundled `http/resources` |
| [`http-file-server`](http-file-server/) | Configurable directory listing and serving | Can expose the selected directory; uses `golog` |
| [`miniweb`](miniweb/) | A configurable static file server | Can expose the selected directory |
| [`http-cookies`](http-cookies/) | Setting, reading, and removing cookies | Demonstration cookie settings only |
| [`http-redirect`](http-redirect/) | Redirect handlers and status codes | Local server |
| [`gorilla-mux-router`](gorilla-mux-router/) | Method-aware routes and path variables | Third-party router; compare modern `ServeMux` patterns |
| [`http-cors`](http-cors/) | CORS response and preflight handling | Third-party router; permissive demo policy |
| [`http-middleware-authentication`](http-middleware-authentication/) | An authentication-shaped middleware | The fixed header check is not real authentication |
| [`http-sessions-gorilla`](http-sessions-gorilla/) | Cookie-backed session state | Hard-coded demo key; never use it in production |
| [`rate-limited-server`](rate-limited-server/) | Per-client token-bucket limiting | Uses `x/time/rate`; documented client map leak |
| [`upload-server`](upload-server/) | Multipart upload parsing | Creates files under `temp_upload`; needs stronger validation and limits |

## Protocol upgrades and advanced examples

| Example | What it teaches | Requirements or caution |
| --- | --- | --- |
| [`gorilla-websockets`](gorilla-websockets/) | Upgrading HTTP and echoing WebSocket messages | Third-party WebSocket package; embedded browser UI |
| [`websocket-chat`](websocket-chat/) | Hub ownership and per-client read/write pumps | Multi-file advanced concurrency example |
| [`https-ssl`](https-ssl/) | Starting a TLS server | Expects old hard-coded certificate paths; modernization required |
| [`http2-client`](http2-client/) | Selecting HTTP/1 or HTTP/2 transports | Expects a local server and `server.crt`; setup is not self-contained |
| [`http2-server`](http2-server/) | TLS server configuration and request inspection | Expects old certificate/GOPATH layout; modernization required |
| [`http2-server-routing-gorilla-mux`](http2-server-routing-gorilla-mux/) | HTTP/2, third-party routing, middleware, and shutdown | Advanced historical integration with external certificates |

## Check your understanding

- Write a handler test with `httptest` without opening a real port.
- Return a JSON error with an appropriate status instead of panicking.
- Add client and server timeouts and explain what each one bounds.
- Propagate request cancellation to downstream work.
- Explain why a hard-coded session key or trusted header is not authentication.

Continue with databases and advanced integrations in stage 6 of the
[learning path](../README.md#learning-path).

## Further reading

- [`net/http` package](https://pkg.go.dev/net/http)
- [Writing Web Applications](https://go.dev/doc/articles/wiki/)
- [Routing enhancements in Go 1.22](https://go.dev/blog/routing-enhancements)
- [`httptest` package](https://pkg.go.dev/net/http/httptest)
- [Go concurrency patterns: context](https://go.dev/blog/context)
