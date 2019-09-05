package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

/*
you can test this with :
$ go run json.go

$ curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost:8080/decode
Elon Musk is 48 years old!

$ curl -s http://localhost:8080/encode
{"firstname":"John","lastname":"Doe","age":25}
*/

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			panic(fmt.Sprintf("ERROR when json.NewDecoder : %v", err))
		}

		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		json.NewEncoder(w).Encode(peter)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Sprintf("ERROR when ListenAndServe : %v", err))
	}
}
