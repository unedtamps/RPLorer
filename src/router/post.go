package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/src/handler"
	"github.com/unedtamps/go-backend/src/middleware"
)

func PostRouter(path string, r *gin.Engine, h handler.PostHandlerI) {
	v1 := r.Group(fmt.Sprintf("%s/v1", path))
	{
		v1.POST("/create", middleware.VerifiyJwtToken, h.CreatePostHandler)
		v1.GET("/account", middleware.VerifiyJwtToken, h.GetPostByAccountHandler)
		v1.GET("/detail/:id", middleware.VerifiyJwtToken, h.GetDetailPostByIDHandler)
		v1.POST("/like/:id", middleware.VerifiyJwtToken, h.LikePostHandler)
		v1.POST("/unlike/:id", middleware.VerifiyJwtToken, h.UnLikePostHandler)
		v1.PATCH("/:id", middleware.VerifiyJwtToken, h.UpdatePostCaptionHandler)
		v1.DELETE("/:id", middleware.VerifiyJwtToken, h.UserDeletePostHandler)
		v1.DELETE(
			"/admin/:id",
			middleware.VerifiyJwtToken,
			middleware.IsAdmin,
			h.AdminDeletePostHandler,
		)
		v1.POST("/comment", middleware.VerifiyJwtToken, h.AddNewCommentHandler)
		v1.GET("/comment/:id", middleware.VerifiyJwtToken, h.GetCommentByIdHandler)
		v1.PATCH("/comment", middleware.VerifiyJwtToken, h.EditCommentHandler)
		v1.GET("/comment", middleware.VerifiyJwtToken, h.GetCommentPostHandler)
		v1.DELETE("/comment/:id", middleware.VerifiyJwtToken, h.DeleteCommentHandler)
	}
}
