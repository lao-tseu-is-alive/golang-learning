package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
$ curl -X GET -H "X-Auth: authenticated" -I http://localhost:8080/api/users
HTTP/1.1 200 OK
Date: Wed, 09 Oct 2019 14:43:29 GMT
Content-Length: 83
Content-Type: text/plain; charset=utf-8


$ curl -X GET -I http://localhost:8080/api/users
HTTP/1.1 401 Unauthorized
Date: Wed, 09 Oct 2019 14:41:47 GMT
Content-Length: 0

*/

func main() {

	// Secured API
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", Secure(func(w http.ResponseWriter,
		r *http.Request) {
		io.WriteString(w, `[{"id":"1","login":"ffghi"},
                           {"id":"2","login":"ffghj"}]`)
	}))
	arguments := os.Args
	listenAddr := ":8080"
	//initialize a logger for server messages output
	logger := log.New(os.Stdout, "HTTP_SERVER: ", log.LstdFlags)

	if len(arguments) == 1 {
		logger.Println("No port number provided as parameter, using 8080 as a default!")
	} else {
		listenAddr = ":" + arguments[1]
	}
	logger.Printf("Starting HTTP server on port %s", listenAddr)
	// Default server
	if err := http.ListenAndServe(listenAddr, mux); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", listenAddr, err)
	}

}

func Secure(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sec := r.Header.Get("X-Auth")
		// here we can check for a valid Json Web Token
		if sec != "authenticated" {
			fmt.Printf("WARNING Secure Middleware refused connection for ip [%s]", r.RemoteAddr)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(w, r) // use the handler
	}

}
