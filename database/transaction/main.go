package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/dblib"
	_ "github.com/lib/pq"
)

const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"
const insert = "INSERT INTO post(ID,TITLE,CONTENT) VALUES (4,'Transaction Title','Transaction Content');"

type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	db := dblib.GetDBConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}
	p := Post{}
	// Query in other session/transaction
	if err := db.QueryRow(selOne, 4).Scan(&p.ID,
		&p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// Query within transaction
	if err := tx.QueryRow(selOne, 4).Scan(&p.ID,
		&p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// After commit or rollback the
	// transaction need to recreated.
	tx.Rollback()

}
