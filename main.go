package main

import (
	"github.com/agugaillard/minesweeper/api/router"
	"github.com/agugaillard/minesweeper/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gameRouter := router.GameRouter{GameService: &service.DefaultGameService{}}
	gameRouter.Routes(r)
	_ = r.Run()
}
