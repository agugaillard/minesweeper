package model

import (
	"errors"
)

type Cell struct {
	Mined     bool
	Flag      Flag
	Explored  bool
	NearMines int
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
// Returns an error if the cell is flagged
func (c *Cell) Explore() (bool, error) {
	if c.Flag != None {
		return false, errors.New("can't explore a flagged cell")
	}
	c.Explored = true
	return c.Mined, nil
}

func (c *Cell) SetFlag(flag Flag) {
	c.Flag = flag
}
