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
	r.Static("/static", "./storage")
	r.MaxMultipartMemory = 2 << 20
	router.AccountRouter("/user", r, h.Acc)
	router.PostRouter("/post", r, h.Post)
	return r
}
