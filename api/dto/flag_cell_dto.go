package dto

import "github.com/agugaillard/minesweeper/domain/model"

type FlagCellRequestDto struct {
	model.Position
	Flag model.Flag `json:"flag"`
}
