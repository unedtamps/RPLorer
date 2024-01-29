package handler

import "github.com/unedtamps/go-backend/src/service"

type Handler struct {
	User UserHandler
	Todo TodoHandler
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		User: newUserHandler(s.User),
		Todo: newTodoHandler(s.Todo),
	}
}
