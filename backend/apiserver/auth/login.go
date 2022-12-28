package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seed/backend/internal/log"
)

func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	remember := context.PostForm("remember")
	log.Infof("username:%s, password:%s, remember:%s", username, password, remember)
	context.String(http.StatusOK, "login success")
}
