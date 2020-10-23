package router

import (
	"github.com/agugaillard/minesweeper/api/dto"
	"github.com/agugaillard/minesweeper/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GameRoutes(r *gin.Engine) {
	r.POST("/game", newGameHandler)
}

func newGameHandler(context *gin.Context) {
	var newGameDto dto.NewGameRequestDto
	err := context.BindJSON(&newGameDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payload",
		})
		return
	}
	game, err := model.NewGame(newGameDto.Cols, newGameDto.Rows, newGameDto.Mines, "")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "unexpected error creating the game",
		})
		return
	}
	context.JSON(http.StatusOK, dto.NewGameDto(game))
}
