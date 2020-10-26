package model

import (
	"strconv"
	"testing"
)

func TestBoardProperties(t *testing.T) {
	expectedCols := 5
	expectedRows := 6
	expectedMines := 10
	board, err := NewRandomBoard(expectedCols, expectedRows, expectedMines)
	if err != nil {
		t.Errorf("unexpected error creating a board")
		return
	}
	if board.Cols != expectedCols {
		t.Errorf("unexpected number of columns")
	}
	if board.Rows != expectedRows {
		t.Errorf("unexpected number of rows")
	}
	mines := 0
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < board.Rows; j++ {
			cell, err := board.getCell(Position{Col: i, Row: j})
			if err != nil {
				t.Errorf("unexpected error counting mines")
				return
			}
			if cell.Mined {
				mines++
			}
		}
	}
	if mines != expectedMines {
		t.Errorf("unexpected number of mines")
	}
}

func TestBoardConstructorValidations(t *testing.T) {
	_, err := NewRandomBoard(5, 5, -1)
	if err == nil {
		t.Errorf("error expected - shouldn't create a board with negative number of mines")
	}
	_, err = NewRandomBoard(5, 5, 0)
	if err == nil {
		t.Errorf("error expected - shouldn't create a board with zero mines")
	}
	_, err = NewRandomBoard(5, 5, 25)
	if err == nil {
		t.Errorf("error expected - shouldn't create a board with all mines")
	}
	_, err = NewRandomBoard(5, 5, 26)
	if err == nil {
		t.Errorf("error expected - shouldn't create a board with more number of mines than cells")
	}
}

func TestBoardGetCell(t *testing.T) {
	expectedCols := 3
	expectedRows := 3
	expectedMines := 2
	board, err := NewRandomBoard(expectedCols, expectedRows, expectedMines)
	if err != nil {
		t.Errorf("unexpected error creating a board")
		return
	}
	_, err = board.getCell(Position{-1, 0})
	if err == nil {
		t.Errorf("expecting error, but none found")
	}
	_, err = board.getCell(Position{0, -1})
	if err == nil {
		t.Errorf("expecting error, but none found")
	}
	_, err = board.getCell(Position{3, 0})
	if err == nil {
		t.Errorf("expecting error, but none found")
	}
	_, err = board.getCell(Position{0, 3})
	if err == nil {
		t.Errorf("expecting error, but none found")
	}
	_, err = board.getCell(Position{1, 1})
	if err != nil {
		t.Errorf("unexpected error, the position is valid")
	}
}

func TestBoardNearMines(t *testing.T) {
	board, err := NewBoard(5, 5, 13, &AllTogetherMinesBoardInitializer{})
	if err != nil {
		t.Errorf("unexpected error creating a board")
		return
	}
	cell, _ := board.getCell(Position{0, 0})
	if cell.NearMines != 3 {
		t.Errorf("Expected 3 near mines, but " + strconv.Itoa(cell.NearMines) + " found")
	}
	cell, _ = board.getCell(Position{0, 1})
	if cell.NearMines != 5 {
		t.Errorf("Expected 5 near mines, but " + strconv.Itoa(cell.NearMines) + " found")
	}
	cell, _ = board.getCell(Position{0, 4})
	if cell.NearMines != 3 {
		t.Errorf("Expected 3 near mines, but " + strconv.Itoa(cell.NearMines) + " found")
	}
	cell, _ = board.getCell(Position{1, 1})
	if cell.NearMines != 8 {
		t.Errorf("Expected 8 near mines, but " + strconv.Itoa(cell.NearMines) + " found")
	}
	cell, _ = board.getCell(Position{1, 2})
	if cell.NearMines != 7 {
		t.Errorf("Expected 7 near mines, but " + strconv.Itoa(cell.NearMines) + " found")
	}
	cell, _ = board.getCell(Position{2, 2})
	if cell.NearMines != 4 {
		t.Errorf("Expected 4 near mines, but " + strconv.Itoa(cell.NearMines) + " found")
	}
	cell, _ = board.getCell(Position{4, 4})
	if cell.NearMines != 0 {
		t.Errorf("Expected 0 near mines, but " + strconv.Itoa(cell.NearMines) + " found")
	}
}

func TestBoardSolved(t *testing.T) {
	board, err := NewBoard(5, 5, 1, &AllTogetherMinesBoardInitializer{})
	if err != nil {
		t.Errorf("unexpected error creating a board")
		return
	}
	for i := board.Cols - 1; i >= 0 && !board.Solved; i-- {
		for j := board.Rows - 1; j >= 0 && !board.Solved; j-- {
			_, err = board.Explore(Position{i, j})
		}
	}
	cell, _ := board.getCell(Position{0, 0})
	if cell.Explored {
		t.Errorf("the board should be Solved before reaching this cell")
	}
	if board.Explored != board.Cols*board.Rows-board.Mines {
		t.Errorf("all cells should be explored when board is Solved")
	}
	if !board.Solved {
		t.Errorf("the board should be Solved")
	}
}

func TestBoardExploreRecursion(t *testing.T) {
	board, err := NewBoard(5, 5, 1, &AllTogetherMinesBoardInitializer{})
	if err != nil {
		t.Errorf("unexpected error creating a board")
		return
	}
	_, _ = board.Explore(Position{4, 4})
	if !board.Solved {
		t.Errorf("the board should be Solved with one exploration")
	}
}

type AllTogetherMinesBoardInitializer struct{}

func (*AllTogetherMinesBoardInitializer) Initialize(board *Board) *Board {
	minesPlaced := 0
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < board.Rows; j++ {
			if minesPlaced < board.Mines {
				board.Cells[i][j] = NewMinedCell()
				minesPlaced++
			} else {
				board.Cells[i][j] = NewSafeCell()
			}
		}
	}
	return board
}
