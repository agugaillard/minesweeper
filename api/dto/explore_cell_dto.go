package dto

import "github.com/agugaillard/minesweeper/domain/model"

type ExploreCellRequestDto struct {
	model.Position
	GameId string `json:"game_id"`
}
