package main

import (
	"github.com/agugaillard/minesweeper/api/router"
	"github.com/agugaillard/minesweeper/data/cache"
	"github.com/agugaillard/minesweeper/data/redis"
	"github.com/agugaillard/minesweeper/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gameRouter := router.GameRouter{
		GameService: service.NewDefaultGameService(cache.GameCache, redis.GameRedis),
	}
	gameRouter.Routes(r)
	_ = r.Run()
}
