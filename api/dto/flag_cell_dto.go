package dto

import "github.com/agugaillard/minesweeper/domain/model"

type FlagCellRequestDto struct {
	model.Position
	Username model.Username `json:"username"`
	Flag     model.Flag     `json:"flag"`
}
