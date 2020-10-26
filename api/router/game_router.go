package router

import (
	"github.com/agugaillard/minesweeper/api/dto"
	apiError "github.com/agugaillard/minesweeper/api/error"
	"github.com/agugaillard/minesweeper/domain/model"
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
	r.PUT("/game/:id/flag", router.flagCellHandler)
	r.POST("/game/:id/save", router.saveHandler)
	r.POST("/game/:id/resume", router.resumeHandler)
}

func (router *GameRouter) newGameHandler(context *gin.Context) {
	var newGameRequest dto.NewGameRequestDto
	err := context.BindJSON(&newGameRequest)
	if ok := handleError(context, err); !ok {
		return
	}
	if newGameRequest.Username == "" {
		handleError(context, apiError.InvalidParameter)
		return
	}
	game, err := router.GameService.NewGame(newGameRequest.Cols, newGameRequest.Rows, newGameRequest.Mines, newGameRequest.Username)
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
	if exploreCellRequest.Username == "" {
		handleError(context, apiError.InvalidParameter)
		return
	}
	err = router.handleAuth(context.Param("id"), exploreCellRequest.Username)
	if ok := handleError(context, err); !ok {
		return
	}
	game, err := router.GameService.ExploreCell(context.Param("id"), exploreCellRequest.Position)
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
	if flagCellRequest.Username == "" {
		handleError(context, apiError.InvalidParameter)
		return
	}
	err = router.handleAuth(context.Param("id"), flagCellRequest.Username)
	if ok := handleError(context, err); !ok {
		return
	}
	err = router.GameService.FlagCell(context.Param("id"), flagCellRequest.Position, flagCellRequest.Flag)
	if ok := handleError(context, err); !ok {
		return
	}
}

func (router *GameRouter) saveHandler(context *gin.Context) {
	var saveGameRequest dto.SaveGameRequestDto
	err := context.BindJSON(&saveGameRequest)
	if ok := handleError(context, err); !ok {
		return
	}
	if saveGameRequest.Username == "" {
		handleError(context, apiError.InvalidParameter)
		return
	}
	err = router.handleAuth(context.Param("id"), saveGameRequest.Username)
	if ok := handleError(context, err); !ok {
		return
	}
	err = router.GameService.Save(context.Param("id"))
	if ok := handleError(context, err); !ok {
		return
	}
}

func (router *GameRouter) resumeHandler(context *gin.Context) {
	var resumeGameRequest dto.ResumeGameRequestDto
	err := context.BindJSON(&resumeGameRequest)
	if ok := handleError(context, err); !ok {
		return
	}
	if resumeGameRequest.Username == "" {
		handleError(context, apiError.InvalidParameter)
		return
	}
	err = router.handleAuth(context.Param("id"), resumeGameRequest.Username)
	if ok := handleError(context, err); !ok {
		return
	}
	game, err := router.GameService.GetGame(context.Param("id"))
	if ok := handleError(context, err); !ok {
		return
	}
	context.JSON(http.StatusOK, dto.NewGameDto(game))
}

func (router *GameRouter) handleAuth(gameId string, username model.Username) error {
	owner, err := router.GameService.GetOwner(gameId)
	if err != nil {
		return err
	}
	if owner != username {
		return apiError.Unauthorized
	}
	return nil
}
