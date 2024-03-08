package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/middleware"
)

func AccountRouter(path string, r *gin.Engine, h handler.AccountHandlerI) {
	v1 := r.Group(fmt.Sprintf("%s/v1", path))
	{
		v1.POST("/register", h.RegisterUserAccount)
		v1.GET("/register/confirm", middleware.VerifiyJwtToken, h.ConfirmRegistrant)
		v1.GET("/register/resend/:email", h.ResendEmailConfirm)
		v1.POST("/login", h.LoginHandler)
	}

	// account := r.Group(path)
	// {
	// 	account.GET("/get", a.GetOne)
	// 	// user.POST("/create", u.CreateUser)
	// 	// user.GET("/getall", m.RateLimit, m.VerifiyJwtToken, u.GetAllUser)
	// 	// user.GET("/me", m.VerifiyJwtToken, u.GetMe)
	// 	// user.POST("/login", u.LoginUser)
	// 	// user.GET("/:email", m.RateLimit, m.VerifiyJwtToken, u.GetUserByEmail)
	// }
}
