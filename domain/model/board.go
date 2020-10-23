package model

import (
	"errors"
)

type Board struct {
	NumCols  int
	NumRows  int
	NumMines int
	Cells    [][]*Cell
	Explored int
	Solved   bool
}

type Position struct {
	Col int `json:"col"`
	Row int `json:"row"`
}

func NewBoard(numCols int, numRows int, numMines int, boardInitializer BoardInitializer) (*Board, error) {
	if numMines <= 0 || numMines >= numCols*numRows {
		return nil, errors.New("invalid number of mines")
	}
	cells := newCellsMatrix(numCols, numRows)
	board := &Board{NumCols: numCols, NumRows: numRows, NumMines: numMines, Cells: cells}
	board = boardInitializer.Initialize(board)
	if err := board.updateNearMines(); err != nil {
		return nil, err
	}
	return board, nil
}

func NewRandomBoard(numCols int, numRows int, numMines int) (*Board, error) {
	return NewBoard(numCols, numRows, numMines, &RandomMinesBoardInitializer{})
}

func newCellsMatrix(numCols int, numRows int) [][]*Cell {
	cells := make([][]*Cell, numCols)
	for i := range cells {
		cells[i] = make([]*Cell, numRows)
	}
	return cells
}

// Complexity: O(m * n)
func (board Board) updateNearMines() error {
	for i := 0; i < board.NumCols; i++ {
		for j := 0; j < board.NumRows; j++ {
			position := Position{Col: i, Row: j}
			nearMines, err := board.countNearMines(position)
			if err != nil {
				return err
			}
			cell, err := board.getCell(position)
			if err != nil {
				return err
			}
			cell.NearMines = nearMines
		}
	}
	return nil
}

// Complexity: O(1)
func (board Board) getCell(position Position) (*Cell, error) {
	if position.Col >= 0 && position.Col < board.NumCols && position.Row >= 0 && position.Row < board.NumRows {
		return board.Cells[position.Col][position.Row], nil
	}
	return nil, errors.New("invalid position")
}

// Returns up to 8 positions
// Complexity: O(1)
func (board Board) getAdjacencies(position Position) []Position {
	adjacencies := make([]Position, 0, 8)
	for i := position.Col - 1; i <= position.Col+1; i++ {
		for j := position.Row - 1; j <= position.Row+1; j++ {
			// if its a valid column
			if i >= 0 && i <= board.NumCols-1 {
				// if its a valid Row
				if j >= 0 && j <= board.NumRows-1 {
					// if its not the cell itself
					if !(i == position.Col && j == position.Row) {
						adjacencies = append(adjacencies, Position{Col: i, Row: j})
					}
				}
			}
		}
	}
	return adjacencies
}

// Returns true if game should finish
func (board *Board) Explore(position Position) (bool, error) {
	currentCell, err := board.getCell(position)
	if err != nil {
		return false, err
	}
	if currentCell.Explored {
		return false, nil
	}
	lost, err := currentCell.Explore()
	if err != nil {
		return false, err
	}
	if lost {
		return lost, nil
	}
	board.Explored++
	if board.Explored == board.NumCols*board.NumRows-board.NumMines {
		board.Solved = true
		return true, nil
	}

	// If there are no mines near it, it should explore all adjacent cells
	if currentCell.NearMines == 0 {
		unflaggedNearPositions, err := board.filterUnflagged(board.getAdjacencies(position))
		if err != nil {
			return false, err
		}
		for _, unflagged := range unflaggedNearPositions {
			_, _ = board.Explore(unflagged)
		}
	}
	return false, nil
}

// Complexity: O(1)
func (board Board) filterUnflagged(adjacencies []Position) ([]Position, error) {
	unflaggedAdjacencies := make([]Position, 0, len(adjacencies))
	for _, adjacency := range adjacencies {
		cell, err := board.getCell(adjacency)
		if err != nil {
			return nil, errors.New("unexpected error")
		}
		if cell.Flag == None {
			unflaggedAdjacencies = append(unflaggedAdjacencies, adjacency)
		}
	}
	return unflaggedAdjacencies, nil
}

// Complexity: O(1)
func (board *Board) countNearMines(position Position) (int, error) {
	adjacencies := board.getAdjacencies(position)
	n := 0
	for _, adjacent := range adjacencies {
		cell, err := board.getCell(adjacent)
		if err != nil {
			return -1, errors.New("unexpected error")
		}
		if cell.Mined {
			n++
		}
	}
	return n, nil
}
