package common

import (
	"booking-room-app/shared/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCreateResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, &model.SingleResponse{

		Status: model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}
