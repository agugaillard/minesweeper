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
	r.POST("/game/:id/flag", router.flagCellHandler)
}

func (router *GameRouter) newGameHandler(context *gin.Context) {
	var newGameDto dto.NewGameRequestDto
	err := context.BindJSON(&newGameDto)
	if ok := handleError(context, err); !ok {
		return
	}
	game, err := router.GameService.NewGame(newGameDto.Cols, newGameDto.Rows, newGameDto.Mines, "")
	if ok := handleError(context, err); !ok {
		return
	}
	context.JSON(http.StatusOK, dto.NewGameDto(game))
}

func (router *GameRouter) exploreCellHandler(context *gin.Context) {
	var exploreCellRequest dto.ExploreCellRequestDto
	err := context.BindJSON(&exploreCellRequest)
	if ok := handleError(context, err); !ok {
		return
	}
	game, err := router.GameService.ExploreCell(exploreCellRequest.GameId, exploreCellRequest.Position)
	if ok := handleError(context, err); !ok {
		return
	}
	context.JSON(http.StatusOK, dto.NewGameDto(game))
}

func (router *GameRouter) flagCellHandler(context *gin.Context) {
	var flagCellRequest dto.FlagCellRequestDto
	err := context.BindJSON(&flagCellRequest)
	if ok := handleError(context, err); !ok {
		return
	}
	err = router.GameService.FlagCell(flagCellRequest.GameId, flagCellRequest.Position, flagCellRequest.Flag)
	if ok := handleError(context, err); !ok {
		return
	}
}
