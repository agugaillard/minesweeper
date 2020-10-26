package dto

import "github.com/agugaillard/minesweeper/domain/model"

type ResumeGameRequestDto struct {
	Username model.Username `json:"username"`
}
