package config

import (
	"github.com/spf13/viper"
	"github.com/unedtamps/go-backend/util"
)

type configServer struct {
	ServerHost       string  `mapstructure:"SERVER_HOST"`
	ServerPort       string  `mapstructure:"SERVER_PORT"`
	JWTSecret        string  `mapstructure:"JWT_SECRET"`
	DBDriver         string  `mapstructure:"DB_DRIVER"`
	PostgresUser     string  `mapstructure:"POSTGRES_USER"`
	PostgresPassword string  `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string  `mapstructure:"POSTGRES_DB"`
	PosgresHost      string  `mapstructure:"POSTGRES_HOST"`
	RedisHost        string  `mapstructure:"REDIS_HOST"`
	RedisPort        string  `mapstructure:"REDIS_PORT"`
	RedisPassword    string  `mapstructure:"REDIS_PASSWORD"`
	RedisDB          int     `mapstructure:"REDIS_DB"`
	Rps              float64 `mapstructure:"LIMIT_RPS"`
	Burst            int     `mapstructure:"LIMIT_BURST"`
	Enable           bool    `mapstructure:"LIMIT_ENABLE"`
	SmtpHost         string  `mapstructure:"SMTP_HOST"`
	SmtPort          int     `mapstructure:"SMTP_PORT"`
	SmtUserName      string  `mapstructure:"SMTP_USERNAME"`
	SmtPassword      string  `mapstructure:"SMTP_PASSWORD"`
	EmailSender      string  `mapstructure:"EMAIL_SENDER"`
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
