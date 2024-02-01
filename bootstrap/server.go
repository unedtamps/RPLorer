package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/config"
	"github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/src"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/service"
)

type Server struct {
	route *gin.Engine
}

func InitServer() (*Server, error) {
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
	return s.route.Run(fmt.Sprintf(":%s", config.Env.Port))
}
