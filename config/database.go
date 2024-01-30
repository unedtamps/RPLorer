package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type dbConfig struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbUri    string `mapstructure:"DB_URI"`
}

func newDBConfig() *dbConfig {
	var conf dbConfig
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}
	return &conf
}

func connectDB() (*sql.DB, error) {
	dbConf := newDBConfig()
	conn, err := sql.Open(dbConf.DbDriver, dbConf.DbUri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
