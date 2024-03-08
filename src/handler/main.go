package handler

import "github.com/unedtamps/go-backend/src/service"

type Handler struct {
	Acc  AccountHandlerI
	Post PostHandlerI
}

type TokenJwt struct {
	Token string `json:"token"`
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Acc:  newAccountHandler(s.Account),
		Post: newPostHandler(s.Post),
	}
}

type paginateForm struct {
	Page      int64 `form:"page"      binding:"number"`
	Page_size int64 `form:"page_size" binding:"number"`
}

func newPageiante(p *paginateForm) (int64, int64) {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Page_size == 0 {
		p.Page_size = 10
	}
	return p.Page_size, (p.Page - 1) * p.Page_size
}
