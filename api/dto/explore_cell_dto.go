package dto

import "github.com/agugaillard/minesweeper/domain/model"

type ExploreCellRequestDto struct {
	model.Position
	Username model.Username `json:"username"`
}
