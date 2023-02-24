package main

import (
	"database/sql"
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/dblib"
	_ "github.com/lib/pq"
)

const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"

/*
*
Sometimes the result of the query or the structure of the table is not clear,
and the result needs to be extracted to some flexible structure.
This brings us to this recipe, where the extraction of values mapped
to column names will be presented.
*/
func main() {
	db := dblib.GetDBConnection()
	defer db.Close()

	rows, err := db.Query(selOne, 1)
	if err != nil {
		panic(err)
	}
	cols, _ := rows.Columns()
	for rows.Next() {
		m := parseWithRawBytes(rows, cols)
		fmt.Println(m)
		m = parseToMap(rows, cols)
		fmt.Println(m)
	}
}

/*
parseWithRawBytes uses the preferred approach, but it is highly dependent
on the driver implementation. It works the way that the slice of RawBytes,
with the same length as the number of the columns in the result, is created.
Because the Scan function requires pointers to values, we need to create
the slice of pointers to the slice of RawBytes (slice of byte slices),
then it can be passed to the Scan function.
After it is successfully extracted, we just remap the values.
In the example code, we cast it to the string because the driver uses
the string type to store the values if the RawBytes is the target.
Beware that the form of stored values depends on driver implementation.
*/

func parseWithRawBytes(rows *sql.Rows, cols []string) map[string]interface{} {
	vals := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(vals))
	for i := range vals {
		scanArgs[i] = &vals[i]
	}
	if err := rows.Scan(scanArgs...); err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	for i, col := range vals {
		if col == nil {
			m[cols[i]] = nil
		} else {
			m[cols[i]] = string(col)
		}
	}
	return m
}

/*
The second approach, parseToMap, is usable in the case that the first one does not work.
It uses almost the same approach, but the slice of values is defined as
the slice of empty interfaces.
This approach relies on the driver.
The driver should determine the default type to assign to the value pointer.
*/
func parseToMap(rows *sql.Rows, cols []string) map[string]interface{} {
	values := make([]interface{}, len(cols))
	pointers := make([]interface{}, len(cols))
	for i := range values {
		pointers[i] = &values[i]
	}

	if err := rows.Scan(pointers...); err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	for i, colName := range cols {
		if values[i] == nil {
			m[colName] = nil
		} else {
			m[colName] = values[i]
		}
	}
	return m
}
