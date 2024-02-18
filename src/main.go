package src

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/router"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	router.UserRouter("/user", r, h.User)
	router.TodoRouter("/todo", r, h.Todo)
	return r
}
