package router

import "github.com/gin-gonic/gin"

type HealthCheckRouter struct {
}

func (router *HealthCheckRouter) Routes(r *gin.Engine) {
	r.GET("/health-check")
}
