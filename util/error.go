package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(err error) interface{} {
	return Response{
		Success: false,
		Error:   err.Error(),
	}
}

func UnknownError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ErrorHandler(err))
}

func LimitError(c *gin.Context, err error) {
	c.JSON(http.StatusTooManyRequests, ErrorHandler(err))
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorHandler(err))
}

func UnauthorizedError(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, ErrorHandler(err))
}

func NotFoundError(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, ErrorHandler(err))
}

func ForbiddenError(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, ErrorHandler(err))
}
