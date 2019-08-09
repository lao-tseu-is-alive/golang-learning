package main

import (
	"database/sql"
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/dblib"
	"github.com/lao-tseu-is-alive/golog"
	_ "github.com/lib/pq"
	"time"
)

const sel = `SELECT id,title,is_done, date_last_modification, id_last_modifier FROM todo;`

const selOne = "SELECT id,title,is_done FROM todo WHERE ID = $1;"

type Todo struct {
	id              int64
	title           string // sql.NullString from "database/sql"
	isDone          bool
	dateLastModif   time.Time
	idLastModifUSer sql.NullInt64
}

func main() {
	db := dblib.GetDBConnection()
	defer db.Close()

	rs, err := db.Query(sel)
	if err != nil {
		golog.Err("Problem in db.Query error: %v", err)
		panic(err)
	}
	defer rs.Close()

	var todos []Todo
	for rs.Next() {
		if rs.Err() != nil {
			panic(rs.Err())
		}
		t := Todo{}
		if err := rs.Scan(&t.id, &t.title, &t.isDone, &t.dateLastModif, &t.idLastModifUSer); err != nil {
			golog.Err("Problem in rs.Scan error: %v", err)
			panic(err)
		}
		todos = append(todos, t)
	}

	var num int
	if rs.NextResultSet() {
		for rs.Next() {
			if rs.Err() != nil {
				panic(rs.Err())
			}
			rs.Scan(&num)
		}
	}

	fmt.Printf("Retrieved posts: %+v\n", todos)
	var count int
	for i, record := range todos {
		fmt.Printf("[%d]\t%s\t%v\n", record.id, record.title, record.isDone)
		count = i + 1
	}

	fmt.Printf("Retrieved number: %d\n", count)

	row := db.QueryRow(selOne, 3)
	todo := Todo{}
	if err := row.Scan(&todo.id, &todo.title, &todo.isDone); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Retrieved one post: %+v\n", todo)

}
