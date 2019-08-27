package main

import (
	"encoding/json"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id      int `json:"id"`
	Title   string
	Authors string
	Isbn10  string
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	golog.Info("GetBook %s", id)
	book := Book{
		Id:      1,
		Title:   "Hamlet",
		Authors: "Shakspeare",
		Isbn10:  "01",
	}
	golog.Info("Book : %v \n", book)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		golog.Err("problem doing Json Encoding : %v", err)
	}
	fmt.Printf("GetBook id: %s\n", id)
}

func NewBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "NewBook id: %s\n", id)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", NewBook).Methods("POST")
	fmt.Println("Navigate to : http://localhost:8080/books/21")
	http.ListenAndServe(":8080", r)
}
