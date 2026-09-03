package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/dblib"
	_ "github.com/lib/pq"
)

const sel = "SELECT * FROM todo p"

func main() {

	db := dblib.GetDBConnection()
	defer db.Close()

	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()
	columns, err := rs.Columns()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected columns: %v\n", columns)

	colTypes, err := rs.ColumnTypes()
	if err != nil {
		panic(err)
	}
	for _, col := range colTypes {
		fmt.Println()
		fmt.Printf("%+v\n", col)
	}

}
