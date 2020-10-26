package dto

import "github.com/agugaillard/minesweeper/domain/model"

type SaveGameRequestDto struct {
	Username model.Username `json:"username"`
}
