package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func APIResponse(ctx *gin.Context, message string, statusCode int, status bool, data interface{}, errors interface{}) {
	jsonResponse := Response{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
	if !status {
		ctx.JSON(statusCode, jsonResponse)
		defer ctx.AbortWithStatus(statusCode)
	} else {
		ctx.JSON(statusCode, jsonResponse)
	}
}

func ValidationResponse(ctx *gin.Context, message string, errors interface{}) {
	APIResponse(ctx, message, http.StatusUnprocessableEntity, false, nil, errors)
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}) {
	APIResponse(ctx, message, http.StatusOK, true, data, nil)
}

func ErrorResponse(ctx *gin.Context, message string, statusCode int, errors interface{}) {
	APIResponse(ctx, message, statusCode, false, nil, errors)
}

func SplitError(err error) []string {
	return strings.Split(err.Error(), "\n")
}
