package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTopic(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello")
}
