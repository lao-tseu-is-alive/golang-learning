package main

import (
	"database/sql"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
	_ "github.com/lib/pq"
)

func main() {
	connStr, err := goutils.GetDbConnectionString("golangdb")
	if err != nil {
		golog.Err("ERROR getting connection string  to DB : %v", err)
		panic(err)
	}
	golog.Info("Connection string : [%s]", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		golog.Err("ERROR connecting to DB : %v", err)
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		golog.Err("ERROR Pinging the DB : %v", err)
		panic(err)
	}
	fmt.Println("DB Ping OK !")
}
