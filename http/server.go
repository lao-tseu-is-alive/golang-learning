package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// example : https://golang.org/src/net/http/example_test.go
	const defaultHost = "localhost"
	const defaultPort = 8080
	port := defaultPort
	val, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		flag.IntVar(&port, "port", defaultPort, "port the server will listen to")
	} else {
		golog.Info("Using ENV variable WEB_PORT to listen %s ", val)
		port, _ = strconv.Atoi(val)
	}

	golog.Info("\nlistening on port : %d\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		golog.Info("connection from %s", r.RemoteAddr)
		fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
		//Iterate over all header fields
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header field :  %q, Value %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
		//Get value for a specified token
		fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])

		fmt.Fprintf(w, "\n\nHello you, %q", html.EscapeString(r.URL.Path))

	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", defaultHost, port), nil))
}
