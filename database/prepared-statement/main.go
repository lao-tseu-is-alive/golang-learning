package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/dblib"
	"github.com/lao-tseu-is-alive/golog"
	_ "github.com/lib/pq"
)

const sel = "SELECT * FROM todo;"
const clean = "DELETE FROM todo WHERE id_creator>2;"
const ins = "INSERT INTO todo (title, id_creator) VALUES ($1, $2);"

var testTable = []struct {
	title     string
	idCreator int
}{
	{"John TODO number One", 3},
	{"Isabella TODO number Two", 4},
	{"John TODO number Two", 3},
}

func main() {
	db := dblib.GetDBConnection()
	defer db.Close()

	// removing from table previous records
	_, err := db.Exec(clean)
	if err != nil {
		golog.Err("ERROR cleaning previous records in DB : %v", err)
		panic(err)
	}

	stm, err := db.Prepare(ins)
	if err != nil {
		golog.Err("ERROR preparing insert statement for DB : %v", err)
		panic(err)
	}

	inserted := int64(0)
	for _, val := range testTable {
		fmt.Printf("Inserting record title: %s\n", val.title)
		// Execute the prepared statement
		r, err := stm.Exec(val.title, val.idCreator)
		if err != nil {
			golog.Err("No way to insert title : [%v]", val.title)
			fmt.Printf("Cannot insert record title : %s\n", val.title)
		}
		affected, err := r.RowsAffected()
		if err != nil {
			golog.Err("No way to get rows affected while insert title : [%v]", val.title)
		} else {
			inserted = inserted + affected
		}
	}

	fmt.Printf("Result: Inserted %d rows.\n", inserted)

}
