package main

import (
	"context"
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/dblib"
	_ "github.com/lib/pq"
	"time"
)

const sel = `SELECT * FROM todo p CROSS JOIN
(SELECT 1 FROM generate_series(1,1000000)) tbl
`

func main() {
	db := dblib.GetDBConnection()
	defer db.Close()

	ctx, canc := context.WithTimeout(context.Background(),
		20*time.Microsecond)
	rows, err := db.QueryContext(ctx, sel)
	canc() //cancel the query
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			continue
		}
		count++
	}

	fmt.Printf("%d rows returned\n", count)

}
