package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		msg := fmt.Sprintf("ERROR when ListenAndServe : %v", err)
		golog.Err(msg)
		log.Println(msg)
		os.Exit(1)
	}
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}
