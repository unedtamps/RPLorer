package config

import (
	"github.com/spf13/viper"
	"github.com/unedtamps/go-backend/util"
)

type configServer struct {
	Host      string `mapstructure:"SERVER_HOST"`
	Port      string `mapstructure:"SERVER_PORT"`
	DBDriver  string `mapstructure:"DB_DRIVER"`
	DBUri     string `mapstructure:"DB_URI"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
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
