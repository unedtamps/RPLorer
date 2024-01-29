package service

import r "github.com/unedtamps/go-backend/internal/repository"

type TodoService struct {
	*r.Store
}

type TodoServiceI interface {
	CreateTodo() error
}

func newTodoService(store *r.Store) *TodoService {
	return &TodoService{
		Store: store,
	}
}

func (t *TodoService) CreateTodo() error {
	return nil
}
