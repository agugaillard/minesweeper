package model

import (
	modelError "github.com/agugaillard/minesweeper/domain/error"
)

type Cell struct {
	Mined     bool `json:"mined"`
	Flag      Flag `json:"flag"`
	Explored  bool `json:"explored"`
	NearMines int  `json:"near_mines"`
}

func NewMinedCell() *Cell {
	cell := newCell()
	cell.Mined = true
	return cell
}

func NewSafeCell() *Cell {
	return newCell()
}

func newCell() *Cell {
	return &Cell{Flag: None}
}

// Returns true if the game should finish
// Errors: ExploreFlagged
func (c *Cell) Explore() (bool, error) {
	if c.Flag != None {
		return false, modelError.ExploreFlagged
	}
	c.Explored = true
	return c.Mined, nil
}

func (c *Cell) SetFlag(flag Flag) {
	c.Flag = flag
}
