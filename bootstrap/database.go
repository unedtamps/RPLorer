package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/unedtamps/go-backend/config"
)

func ConnectDB() (*sql.DB, error) {
	db_url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.PosgresHost,
		config.Env.PostgresUser,
		config.Env.PostgresPassword,
		config.Env.PostgresDB,
	)
	conn, err := sql.Open(config.Env.DBDriver, db_url)
	if err != nil {
		return nil, err
	}
	return conn, conn.Ping()
}
