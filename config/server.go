package config

import (
	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/src"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/service"
)

var (
	host = "localhost"
	port = 8080
)

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
	return s.route.Run(":8080")
}
