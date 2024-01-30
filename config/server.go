package config

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/src"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/service"
)

var (
	host = "localhost"
	port = 8080
)

type serverConfig struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port string `mapstructure:"SERVER_PORT"`
}

func newServerConfig() *serverConfig {
	var conf serverConfig
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

type Server struct {
	route *gin.Engine
}

func NewServer() (*Server, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	repo := repository.NewStore(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	router := src.NewRouter(handler)
	return &Server{router}, nil
}

func (s *Server) Start() error {
	server := newServerConfig()
	return s.route.Run(fmt.Sprintf(":%s", server.Port))
}
