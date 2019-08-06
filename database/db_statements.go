package main

import (
	"database/sql"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
	_ "github.com/lib/pq"
)

const sel = "SELECT * FROM todo;"
const trunc = "DELETE FROM todo WHERE id_creator=2;"
const ins = `INSERT INTO todo (title, id_creator) VALUES 
                                ('practice more golang', 2), 
                                ('experiment with golang concurrency', 2);
`

func main() {
	db := createConnection()
	defer db.Close()

	r, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}
	affected, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted %v records.\n", affected)
	r, err = db.Exec(ins)
	if err != nil {
		panic(err)
	}
	affected, err = r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted %d records !\n", affected)

	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	count := 0
	for rs.Next() {
		println(rs)
		count++
	}
	fmt.Printf("Total of %d records selected.\n", count)
}

func createConnection() *sql.DB {
	connStr, err := goutils.GetDbConnectionString("golangdb")
	if err != nil {
		golog.Err("ERROR getting connection string  to DB : %v", err)
		panic(err)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		golog.Err("ERROR opening connection to DB : %v", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		golog.Err("ERROR doing ping to DB : %v", err)
		panic(err)
	}
	return db
}
