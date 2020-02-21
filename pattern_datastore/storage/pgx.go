package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
	"github.com/lao-tseu-is-alive/golog"
)

const (
	getPGVersion = "SELECT version();"
)

// PGX struct
type PGX struct {
	Conn *pgxpool.Pool
}

func NewPgxDB(dbConnectionString string, maxConnectionsInPool int) (DB, error) {
	var psql PGX
	var successOrFailure string = "OK"

	var parsedConfig *pgx.ConnConfig
	var err error
	parsedConfig, err = pgx.ParseConfig(dbConnectionString)
	if err != nil {
		return nil, err
	}

	dbHost := parsedConfig.Host
	dbPort := parsedConfig.Port
	dbUser := parsedConfig.User
	dbPass := parsedConfig.Password
	dbName := parsedConfig.Database

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s pool_max_conns=%d", dbHost, dbPort, dbUser, dbPass, dbName, maxConnectionsInPool)

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println(dsn)
	connPool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		successOrFailure = "FAILED"
		golog.Info("Connecting to database %s as user %s : %s \n", dbName, dbUser, successOrFailure)
		golog.Fatal("ERROR TRYING DB CONNECTION : %v ", err)
	} else {
		golog.Info("Connecting to database %s as user %s : %s \n", dbName, dbUser, successOrFailure)
		golog.Info("Fetching one record to test if db connection is valid...\n")
		var version string
		if errPing := connPool.QueryRow(context.Background(), getPGVersion).Scan(&version); errPing != nil {
			golog.Err("Connection is invalid ! ")
			golog.Fatal("DB ERROR scanning row: %s", errPing)
		}
		golog.Info("SUCCESS Connecting to Postgres version : [%s]", version)
	}

	fmt.Println("--------------------------------------------------------------------------------------------")
	psql.Conn = connPool
	return &psql, err
}

func (db *PGX) New(val Todo) (string, error) {
	const todoInsert = "INSERT INTO todo (title, is_done) VALUES($1,$2) RETURNING id"
	lastInsertId, err := db.getQueryRowInt(todoInsert, val.Title, val.IsDone)
	if err != nil {
		golog.Err("[1013] user could not be created in DB. failed storage.ExecActionQuery(val:%v) %v", val, err)
		return "", ErrCouldNotBeCreated
	}
	if lastInsertId < 1 {
		return "", ErrCouldNotBeCreated
		//return echo.NewHTTPError(http.StatusInternalServerError, "user was not created in DB (no rowsAffected)")
	}
	newTodo := Todo{
		ID:     lastInsertId,
		Title:  val.Title,
		IsDone: val.IsDone,
	}
	res, _ := json.Marshal(newTodo)
	return string(res), nil
}

func (db *PGX) Get(key int) (string, error) {
	const todoGetJson = `SELECT row_to_json(u) FROM (
	SELECT id, title, is_done FROM todo WHERE id=$1) As u;`
	jsonResult, err := db.getQueryRowString(todoGetJson, key)
	if err != nil {
		return "", ErrNotFound
	}
	return jsonResult, nil
}

func (db *PGX) List() (string, error) {
	const todoListJson = `SELECT json_agg(row_to_json(u)) FROM  (
	SELECT id, title, is_done FROM todo ORDER BY id) As u;`
	jsonResult, err := db.getQueryRowString(todoListJson)
	if err != nil {
		return "", err
	}
	return jsonResult, nil
}

func (db *PGX) getQueryRowInt(sql string, arguments ...interface{}) (result int, err error) {
	err = db.Conn.QueryRow(context.Background(), sql, arguments...).Scan(&result)
	if err != nil {
		golog.Err("Exec unexpectedly failed with %v: %v", sql, err)
		return 0, err
	}
	return result, err
}

func (db *PGX) getQueryRowBool(sql string, arguments ...interface{}) (result bool, err error) {
	err = db.Conn.QueryRow(context.Background(), sql, arguments...).Scan(&result)
	if err != nil {
		golog.Err("Exec unexpectedly failed with %v: %v", sql, err)
		return false, err
	}
	return result, err
}

func (db *PGX) getQueryRowString(sql string, arguments ...interface{}) (result string, err error) {
	err = db.Conn.QueryRow(context.Background(), sql, arguments...).Scan(&result)
	if err != nil {
		golog.Err("Exec unexpectedly failed with %v: %v", sql, err)
		return "", err
	}
	return result, err
}
