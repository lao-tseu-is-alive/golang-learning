package main

import (
	"log"
	"net/http"
	"os"
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
	err := http.ListenAndServe(listenAddr, http.FileServer(http.Dir(`.`)))
	if err != nil {
		logger.Fatalf("Unable to run webserver ! Error : %v", err)
	}

}
