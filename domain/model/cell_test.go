package model

import (
	"testing"
)

func TestCellProperties(t *testing.T) {
	board, err := NewBoard(5, 5, 13, &AllTogetherMinesBoardInitializer{})
	if err != nil {
		t.Errorf("unexpected error creating a board")
		return
	}
	cell, _ := board.getCell(Position{0, 0})
	if cell.Explored {
		t.Errorf("cell shouldn't be explored")
	}
	if cell.Flag != None {
		t.Errorf("cell shouldn't be flagged")
	}
}

func TestCellExplore(t *testing.T) {
	board, err := NewBoard(5, 5, 13, &AllTogetherMinesBoardInitializer{})
	if err != nil {
		t.Errorf("unexpected error creating a board")
		return
	}
	cell, _ := board.getCell(Position{0, 0})
	finish, err := cell.Explore()
	if err != nil {
		t.Errorf("unexpected error exploring a cell")
	}
	if !finish {
		t.Errorf("game should have ended")
	}
	cell, _ = board.getCell(Position{4, 4})
	finish, err = cell.Explore()
	if err != nil {
		t.Errorf("unexpected error exploring a cell")
	}
	if finish {
		t.Errorf("game shouldn't have ended")
	}
	if !cell.Explored {
		t.Errorf("cell should be explored")
	}
	cell, _ = board.getCell(Position{4, 3})
	cell.Flag = RedFlag
	_, err = cell.Explore()
	if err == nil {
		t.Errorf("error expected - shouldn't explore a flagged cell")
	}
	cell.Flag = QuestionMark
	_, err = cell.Explore()
	if err == nil {
		t.Errorf("error expected - shouldn't explore a flagged cell")
	}
}
