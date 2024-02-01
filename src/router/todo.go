package router

import (
	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/middleware"
)

func TodoRouter(path string, r *gin.Engine, t handler.TodoHandlerI) {
	todo := r.Group(path)
	{
		todo.POST("/create", middleware.VerifiyJwtToken, t.CreateTodo)
		todo.GET("/", middleware.VerifiyJwtToken, t.GetTodoByUserId)
	}
}
