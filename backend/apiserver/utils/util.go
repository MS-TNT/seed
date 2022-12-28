package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"seed/backend/cfg"
	"strconv"
)

func GetIdFromUrl(key string, context *gin.Context) (int, error) {
	if id, err := strconv.Atoi(context.Param(key)); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    cfg.DataSourceParamError,
			"message": fmt.Sprintf("%s id is illegal. please check param", key),
		})
		return -1, err
	} else {
		return id, nil
	}
}
