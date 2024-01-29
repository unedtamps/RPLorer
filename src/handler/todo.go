package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/service"
)

type TodoHandler struct {
	t service.TodoServiceI
}

type TodoHandlerI interface {
	CreateTodo(*gin.Context)
}

func newTodoHandler(todoService service.TodoServiceI) TodoHandler {
	return TodoHandler{t: todoService}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hi": "berhasil",
	})
}
