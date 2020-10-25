package main

import (
	"github.com/agugaillard/minesweeper/api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.GameRoutes(r)
	_ = r.Run()
}
