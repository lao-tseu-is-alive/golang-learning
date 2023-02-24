package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr, err := goutils.GetDbConnectionString("golangdb")
	if err != nil {
		panic(err)
	}
	golog.Info("Connection string : [%s]", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping OK.")
	ctx, _ := context.WithTimeout(context.Background(),
		time.Nanosecond)
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	// Verify the connection is
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = conn.PingContext(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Ping OK.")

}
