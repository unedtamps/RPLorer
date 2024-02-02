package util

import (
	"fmt"
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
	c.JSON(http.StatusInternalServerError, ErrorHandler(fmt.Errorf("Unknown Error: %v", err)))
}

func LimitError(c *gin.Context, err error) {
	c.JSON(http.StatusTooManyRequests, ErrorHandler(fmt.Errorf("Limit Error: %v", err)))
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorHandler(fmt.Errorf("Bad Request: %v", err)))
}

func UnauthorizedError(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, ErrorHandler(fmt.Errorf("Unauthorized: %v", err)))
}

func NotFoundError(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, ErrorHandler(fmt.Errorf("Not Found: %v", err)))
}

func ForbiddenError(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, ErrorHandler(fmt.Errorf("Forbidden :%v", err)))
}
