package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GameRoutes(r *gin.Engine) {
	r.GET("/ping", pingpong)
}

func pingpong(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}
