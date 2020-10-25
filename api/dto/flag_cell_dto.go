package dto

import "github.com/agugaillard/minesweeper/domain/model"

type FlagCellRequestDto struct {
	model.Position
	GameId string `json:"game_id"`
	Flag model.Flag `json:"flag"`
}
