package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const ReadTimeoutDefault = 5 * time.Second
const WriteTimeoutDefault = 10 * time.Second
const IdleTimeoutDefault = 15 * time.Second

func main() {
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"time\":\"%s\"}", now.Format(time.RFC3339Nano))
	})
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      nil,
		ErrorLog:     logger,
		ReadTimeout:  ReadTimeoutDefault,
		WriteTimeout: WriteTimeoutDefault,
		IdleTimeout:  IdleTimeoutDefault,
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", listenAddr, err)
	}
}
