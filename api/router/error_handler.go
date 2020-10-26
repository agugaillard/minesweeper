package router

import (
	apiError "github.com/agugaillard/minesweeper/api/error"
	dataError "github.com/agugaillard/minesweeper/data/error"
	modelError "github.com/agugaillard/minesweeper/domain/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(context *gin.Context, err error) bool {
	if err != nil {
		var status int
		switch err {
		case modelError.InvalidNumberOfMines,
			modelError.InvalidPosition,
			modelError.ExploreFlagged:
			status = http.StatusBadRequest
		case dataError.GameNotFound:
			status = http.StatusNotFound
		case dataError.GameAlreadyExists:
			status = http.StatusInternalServerError
		case apiError.Unauthorized:
			status = http.StatusForbidden
		default:
			status = http.StatusInternalServerError
		}
		context.JSON(status, gin.H{
			"error": err.Error(),
		})
		return false
	}
	return true
}
