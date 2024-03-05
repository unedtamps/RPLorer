package handler

import "github.com/unedtamps/go-backend/src/service"

type Handler struct {
	Acc AccountHandlerI
}

type TokenJwt struct {
	Token string `json:"token"`
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Acc: newAccountHandler(s.Account),
	}
}

type paginateForm struct {
	Page      int64 `form:"page"      binding:"min=1"`
	Page_size int64 `form:"page_size" binding:"min=5"`
}
