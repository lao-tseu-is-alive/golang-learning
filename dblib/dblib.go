package dblib

import (
	"database/sql"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
)

func GetDBConnection() *sql.DB {
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
