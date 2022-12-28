package handlers

import (
	"github.com/gin-gonic/gin"
)

func Abort(context *gin.Context, httpCode, businessCode int, msg string) {
	//TODO
	//need reflect businessCode with httpCode
	context.JSON(httpCode, map[string]interface{}{
		"code":    businessCode,
		"message": msg,
	})
}
