package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/dblib"
	_ "github.com/lib/pq"
)

const call = "select * from format_name($1,$2,$3)"

type Result struct {
	Name     string
	Category int
}

func main() {
	db := dblib.GetDBConnection()
	defer db.Close()
	r := Result{}

	if err := db.QueryRow(call, "John", "Doe", 32).Scan(&r.Name); err != nil {
		panic(err)
	}
	fmt.Printf("Result is: %+v\n", r)
}
