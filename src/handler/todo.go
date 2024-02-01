package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/middleware"
	"github.com/unedtamps/go-backend/src/service"
	"github.com/unedtamps/go-backend/util"
)

type TodoHandler struct {
	t service.TodoServiceI
}

type TodoHandlerI interface {
	CreateTodo(*gin.Context)
	GetTodoByUserId(*gin.Context)
}

type createTodoBody struct {
	Title       string `json:"title"       binding:"required,max=255"`
	Description string `json:"description" binding:"required"`
}

type userParams struct {
	UserId string `uri:"user_id" binding:"required"`
}

func newTodoHandler(todoService service.TodoServiceI) TodoHandlerI {
	return &TodoHandler{todoService}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	body := createTodoBody{}
	cred := middleware.GetCredentials(c)
	if err := c.ShouldBindJSON(&body); err != nil {
		util.BadRequest(c, err)
		return
	}

	todo, err := h.t.CreateTodo(c, service.TodoParams{
		Title:  body.Title,
		Desc:   body.Description,
		UserId: cred.Id,
	})
	if err != nil {
		util.UnknownError(c, err)
		return
	}
	util.ResponseCreated(c, "Todo created", todo)
}

func (h *TodoHandler) GetTodoByUserId(c *gin.Context) {
	cred := middleware.GetCredentials(c)
	query := paginateForm{}
	if err := c.ShouldBindQuery(&query); err != nil {
		util.BadRequest(c, err)
		return
	}
	todo, metadata, err := h.t.GetTodoByUserId(c, cred.Id, query.Page, query.Page_size)
	if err != nil {
		util.UnknownError(c, err)
		return
	}
	if todo == nil {
		util.NotFoundError(c, errors.New("Todo"))
		return
	}
	util.ResponseData(c, "Get todo", metadata, todo)
}
