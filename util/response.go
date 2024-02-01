package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message,omitempty"`
	MetaData *MetaData   `json:"meta_data,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Error    interface{} `json:"error,omitempty"`
}

type MetaData struct {
	Page      int64       `json:"page"`
	TotalData int64       `json:"total_data"`
	Details   interface{} `json:"details,omitempty"`
}

func WithPagination(page int64, page_size int64) (int64, int64) {
	return page_size, (page - 1) * page_size
}

func WithMetadata(page int64, totalData int64, details interface{}) MetaData {
	return MetaData{
		Page:      page,
		TotalData: totalData,
		Details:   details,
	}
}

func responseWithData(message string, data interface{}, metadata *MetaData) interface{} {
	return Response{
		Success:  true,
		Message:  message,
		Data:     data,
		MetaData: metadata,
	}
}

func responseOk(message string) interface{} {
	return Response{
		Success: true,
		Message: message,
	}
}

func ResponseCreated(c *gin.Context, message string, data ...interface{}) {
	c.JSON(http.StatusCreated, responseWithData(message, data, nil))
}

func ResponseData(c *gin.Context, message string, metadata *MetaData, data ...interface{}) {
	c.JSON(http.StatusOK, responseWithData(message, data, metadata))
}

func ResponseOk(c *gin.Context, message string) {
	c.JSON(http.StatusOK, responseOk(message))
}
