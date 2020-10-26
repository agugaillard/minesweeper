package model

import (
	modelError "github.com/agugaillard/minesweeper/domain/error"
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

// Errors: InvalidNumberOfMines
func NewGame(numCols int, numRows int, numMines int, username Username) (*Game, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	board, err := NewRandomBoard(numCols, numRows, numMines)
	return &Game{Id: id.String(), Start: time.Now(), Owner: username, Board: board}, nil
}

// Errors: InvalidPosition, InteractFinished, ExploreFlagged
func (g *Game) Explore(position Position) error {
	if g.Finished {
		return modelError.InteractFinished
	}
	finish, err := g.Board.Explore(position)
	if err != nil {
		return err
	}
	g.Finished = finish || g.Board.Solved
	if g.Finished {
		g.End = time.Now()
	}
	return nil
}

// Errors: InvalidPosition, InteractFinished
func (g *Game) Flag(position Position, flag Flag) error {
	if g.Finished {
		return modelError.InteractFinished
	}
	return g.Board.Flag(position, flag)
}
