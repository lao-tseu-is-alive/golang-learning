package main

import (
	"fmt"
	"net/http"
)

const (
	defaultHost        = "localhost"
	defaultPort        = 8080
	defaultTitle       = "Simple Golang HTTP2 Server"
	defaultDescription = "Basic Golang HTTP2 Server"
	defaultSSLKeyFile  = "/etc/ssl/private/lausanne_ch_2019.key"
	defaultSSLCertFile = "/etc/ssl/certs/lausanne_ch_2019.crt"
)

type SimpleHTTP struct{}

func (s SimpleHTTP) ServeHTTP(rw http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(rw, "<h1>Hello SSL world!</h1>")
}

func main() {
	fmt.Printf("Starting HTTP server on port %0d", defaultPort)
	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}
	if err := s.ListenAndServeTLS(defaultSSLCertFile, defaultSSLKeyFile); err != nil {
		panic(err)
	}
}
