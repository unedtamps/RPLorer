package config

import (
	"github.com/spf13/viper"
	"github.com/unedtamps/go-backend/util"
)

type configServer struct {
	ServerHost       string `mapstructure:"SERVER_HOST"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	DBDriver         string `mapstructure:"DB_DRIVER"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PosgresHost      string `mapstructure:"POSTGRES_HOST"`
	RedisHost        string `mapstructure:"REDIS_HOST"`
	RedisPort        string `mapstructure:"REDIS_PORT"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	RedisDB          int    `mapstructure:"REDIS_DB"`
}

var Env configServer

func init() {
	viper.SetConfigFile(".env")
	Env = configServer{}
	err := viper.ReadInConfig()
	if err != nil {
		util.Log.Fatal(err)
	}
	err = viper.Unmarshal(&Env)
	if err != nil {
		util.Log.Fatal(err)
	}
}
