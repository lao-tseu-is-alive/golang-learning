package main

import (
	"log"
	"net/http"
)

func main() {

	fileSrv := http.FileServer(http.Dir("http/resources"))
	fileSrv = http.StripPrefix("/html", fileSrv)

	http.HandleFunc("/welcome", serveWelcome)
	http.Handle("/html/", fileSrv)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveWelcome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "http/resources/welcome.txt")
}
