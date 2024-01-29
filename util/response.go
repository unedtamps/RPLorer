package util

import "github.com/gin-gonic/gin"

func ErrorHandler(err error) gin.H {
	return gin.H{
		"success": false,
		"error":   err.Error(),
	}
}

func ResponseHandler(data interface{}) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}
