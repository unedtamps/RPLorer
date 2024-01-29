package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/todoapp?sslmode=disable"
)

func connectDB() (*sql.DB, error) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
