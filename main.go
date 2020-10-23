package main

import (
	"github.com/agugaillard/minesweeper/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.GameRoutes(r)
	_ = r.Run()
}
