package storage

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrCouldNotBeCreated = errors.New("could not be created in DB")
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

// DB is the interface for a simple table store.
type DB interface {
	New(val Todo) (string, error)
	Get(key int) (string, error)
	List() (string, error)
}

// InitDB with appropriate driver
func InitDB(dbDriver, dbConnectionString string, maxConnectionCount int) (DB, error) {
	var err error
	var db DB

	if dbDriver == "pgx" {
		db, err = NewPgxDB(dbConnectionString, runtime.NumCPU())
		if err != nil {
			return nil, fmt.Errorf("error opening postgresql database with pgx driver: %s", err)
		}
	} else if dbDriver == "memory" {
		db, err = NewMemoryDB()
		if err != nil {
			return nil, fmt.Errorf("error opening memory store: %s", err)
		}
	} else if dbDriver == "none" {
		db = nil
	} else {
		return nil, errors.New("unsupported DB driver type")
	}

	return db, nil
}
