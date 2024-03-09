package bootstrap

import (
	"crypto/tls"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/config"
	"github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/src"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/service"
	"gopkg.in/gomail.v2"
)

type Server struct {
	route *gin.Engine
}

func InitServer() (*Server, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	cache, err := connectCache()
	if err != nil {
		return nil, err
	}
	// email dialer
	dialer := gomail.NewDialer(
		config.Env.SmtpHost,
		config.Env.SmtPort,
		config.Env.SmtUserName,
		config.Env.SmtPassword,
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	repo := repository.NewStore(db)
	service := service.NewService(repo, cache, dialer)
	handler := handler.NewHandler(service)
	router := src.NewRouter(handler)
	return &Server{router}, nil
}

func (s *Server) Start() error {
	return s.route.Run(fmt.Sprintf(":%s", config.Env.ServerPort))
}
