package bootstrap

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/unedtamps/go-backend/config"
)

func connectDB() (*sql.DB, error) {
	conn, err := sql.Open(config.Env.DBDriver, config.Env.DBUri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
