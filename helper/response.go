package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessJSON(ctx *gin.Context, message string, code int, data interface{}) Response {
	resp := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	return resp
}

func ErrorJSON(ctx *gin.Context, message string, code int, data interface{}) Response {
	resp := Response{
		Status:  "failed",
		Message: message,
		Data:    data,
	}

	return resp
}
