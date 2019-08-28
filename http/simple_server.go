package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><body><h3>The simple Go http server !</h3><img src='/resources/logo.svg'></body></html>")
	})

	fs := http.FileServer(http.Dir("resources/"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Sprintf("ERROR when ListenAndServe : %v", err))
	}
}
