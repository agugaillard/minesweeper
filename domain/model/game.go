package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Game struct {
	Id       string
	Start    time.Time
	End      time.Time
	Owner    Username
	Board    *Board
	Finished bool
}

func NewGame(numCols int, numRows int, numMines int, username Username) (*Game, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("unexpected error creating a game")
	}
	board, err := NewRandomBoard(numCols, numRows, numMines)
	return &Game{Id: id.String(), Start: time.Now(), Owner: username, Board: board}, nil
}

func (g *Game) Explore(position Position) {
	_, _ = g.Board.Explore(position)
}
