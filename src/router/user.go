package router

import (
	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/handler"
	m "github.com/unedtamps/go-backend/src/middleware"
)

func UserRouter(path string, r *gin.Engine, u handler.UserHandlerI) {
	user := r.Group(path)
	{
		user.POST("/create", u.CreateUser)
		user.GET("/getall", m.VerifiyJwtToken, u.GetAllUser)
		user.GET("/me", m.VerifiyJwtToken, u.GetMe)
		user.POST("/login", u.LoginUser)
	}
}
