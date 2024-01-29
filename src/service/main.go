package service

import r "github.com/unedtamps/go-backend/internal/repository"

type Service struct {
	User UserServiceI
	Todo TodoServiceI
}

func NewService(repo *r.Store) *Service {
	return &Service{
		User: newUserService(repo),
		Todo: newTodoService(repo),
	}
}
