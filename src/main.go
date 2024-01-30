package src

import (
	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/handler"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/create-user", h.User.CreateUser)
	router.POST("/create-todo/:user_id", h.Todo.CreateTodo)
	router.GET("/get-todo/:user_id", h.Todo.GetTodoByUserId)
	router.GET("/get-alluser", h.User.GetAllUser)
	return router
}
