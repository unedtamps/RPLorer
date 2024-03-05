package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/handler"
)

func AccountRouter(path string, r *gin.Engine, a handler.AccountHandlerI) {
	v1 := r.Group(fmt.Sprintf("%s/v1", path))
	{
		v1.GET("/get", a.GetOne)
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
