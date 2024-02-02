package service

import (
	"github.com/redis/go-redis/v9"
	r "github.com/unedtamps/go-backend/internal/repository"
)

type Service struct {
	User UserServiceI
	Todo TodoServiceI
}

func NewService(repo *r.Store, cache *redis.Client) *Service {
	return &Service{
		User: newUserService(repo, cache),
		Todo: newTodoService(repo, cache),
	}
}
