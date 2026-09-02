package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
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
	parsedConfig, err := pgxpool.ParseConfig(dbConnectionString)
	if err != nil {
		return nil, fmt.Errorf("parse PostgreSQL configuration: %w", err)
	}
	parsedConfig.MaxConns = int32(maxConnectionsInPool)

	connPool, err := pgxpool.NewWithConfig(context.Background(), parsedConfig)
	if err != nil {
		return nil, fmt.Errorf("create PostgreSQL connection pool: %w", err)
	}

	var version string
	if err := connPool.QueryRow(context.Background(), getPGVersion).Scan(&version); err != nil {
		connPool.Close()
		return nil, fmt.Errorf("verify PostgreSQL connection: %w", err)
	}

	golog.Info(
		"Connected to PostgreSQL database %s on %s:%d as user %s (version: %s)",
		parsedConfig.ConnConfig.Database,
		parsedConfig.ConnConfig.Host,
		parsedConfig.ConnConfig.Port,
		parsedConfig.ConnConfig.User,
		version,
	)
	return &PGX{Conn: connPool}, nil
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
