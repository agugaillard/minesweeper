package router

import (
	"github.com/agugaillard/minesweeper/api/dto"
	"github.com/agugaillard/minesweeper/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GameRouter struct {
	GameService service.GameService
}

func (router *GameRouter) Routes(r *gin.Engine) {
	r.POST("/game", router.newGameHandler)
	r.POST("/game/:id/explore", router.exploreCellHandler)
}

func (router *GameRouter) newGameHandler(context *gin.Context) {
	var newGameDto dto.NewGameRequestDto
	err := context.BindJSON(&newGameDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payload",
		})
		return
	}
	game, err := router.GameService.NewGame(newGameDto.Cols, newGameDto.Rows, newGameDto.Mines, "")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "unexpected error creating the game",
		})
		return
	}
	context.JSON(http.StatusOK, dto.NewGameDto(game))
}

func (router *GameRouter) exploreCellHandler(context *gin.Context) {
	var exploreCellRequest dto.ExploreCellRequestDto
	err := context.BindJSON(&exploreCellRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payload",
		})
		return
	}
	game, err := router.GameService.ExploreCell(exploreCellRequest.GameId, exploreCellRequest.Position)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "game not found",
		})
		return
	}
	context.JSON(http.StatusOK, dto.NewGameDto(game))
}
