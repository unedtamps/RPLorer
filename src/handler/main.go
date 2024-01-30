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

type paginateForm struct {
	Page      int64 `form:"page"      binding:"min=1"`
	Page_size int64 `form:"page_size" binding:"min=5"`
}
