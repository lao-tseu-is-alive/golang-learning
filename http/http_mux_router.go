package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

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

	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "User GET")
		}
		if r.Method == http.MethodPost {
			fmt.Fprintln(w, "User POST")
		}
	})

	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			now := time.Now()
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "{\"time\":\"%s\"}", now.Format(time.RFC3339))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

	})

	// separate handler
	itemMux := http.NewServeMux()
	itemMux.HandleFunc("/items/clothes", func(w http.ResponseWriter,
		r *http.Request) {
		fmt.Fprintln(w, "Clothes")
	})
	mux.Handle("/items/", itemMux)

	// Admin handlers
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/ports", func(w http.ResponseWriter,
		r *http.Request) {
		fmt.Fprintln(w, "Ports")
	})

	mux.Handle("/admin/", http.StripPrefix("/admin",
		adminMux))

	// Default server
	if err := http.ListenAndServe(listenAddr, mux); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", listenAddr, err)
	}
}
