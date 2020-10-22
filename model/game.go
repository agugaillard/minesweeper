package model

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Game struct {
	Id       string    `json:"id"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Owner    Username  `json:"owner"`
	Board    *Board    `json:"board"`
	Finished bool      `json:"finished"`
}

func NewGame(numCols int, numRows int, numMines int, username Username) *Game {
	id, err := uuid.NewRandom()
	if err != nil {
		context.TODO()
		return nil
	}
	return &Game{Id: id.String(), Start: time.Now(), Owner: username, Board: NewRandomBoard(numCols, numRows, numMines)}
}
