package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/pattern_datastore/storage"
	"github.com/lao-tseu-is-alive/golog"
	"runtime"
)

func main() {
	dbDriver := flag.String("db", "pgx", "db connection driver [values: pgx || memory ]")
	dbConnectionString := flag.String("db_connection_string",
		"host=localhost user=gouser password=gouser-2019 dbname=golangdb sslmode=disable",
		"db connection string")
	flag.Parse()
	// init database with appropriate driver
	db, err := storage.InitDB(*dbDriver, *dbConnectionString, runtime.NumCPU()*4)
	if err != nil {
		golog.Fatal("error  calling InitDB : %v ", err)
	}

	t1 := storage.Todo{
		ID:     1,
		Title:  "Learn Golang",
		IsDone: false,
	}
	t2 := storage.Todo{
		ID:     1,
		Title:  "Learn Test Driven Dev",
		IsDone: false,
	}

	type responseNew struct {
		Id int `json:id`
	}

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("Using DRIVER %v\n", *dbDriver)
	fmt.Println("db.New --> val1, err := db.New(t1) ")
	val1, err := db.New(t1)
	fmt.Println("val1      : ", val1)
	res := storage.Todo{}
	json.Unmarshal([]byte(val1), &res)
	get1, err := db.Get(res.ID)
	if err != nil {
		golog.Err("db.Get(%v) failed error: %v", 1, err)
	}
	fmt.Println("db.Get(1) : ", get1)
	fmt.Println("db.New --> val2, err := db.New(t2) ")
	val2, err := db.New(t2)
	fmt.Println("val2      : ", val2)
	res2 := storage.Todo{}
	json.Unmarshal([]byte(val2), &res2)
	get2, err := db.Get(res2.ID)
	if err != nil {
		golog.Err("db.Get(%v) failed error: %v", 2, err)
	}
	fmt.Println("db.Get(2) : ", get2)
	list, err := db.List()
	fmt.Println("db.List() : ", list)
	fmt.Println("--------------------------------------------------------------------------------------------")

}
