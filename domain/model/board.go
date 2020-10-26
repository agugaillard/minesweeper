package model

import (
	modelError "github.com/agugaillard/minesweeper/domain/error"
)

type Board struct {
	Cols     int       `json:"cols"`
	Rows     int       `json:"rows"`
	Mines    int       `json:"mines"`
	Cells    [][]*Cell `json:"cells"`
	Explored int       `json:"explored"`
	Solved   bool      `json:"solved"`
}

type Position struct {
	Col int `json:"col"`
	Row int `json:"row"`
}

// Errors: InvalidBoardProperties
func NewBoard(numCols int, numRows int, numMines int, boardInitializer BoardInitializer) (*Board, error) {
	if numMines <= 0 || numMines >= numCols*numRows {
		return nil, modelError.InvalidBoardProperties
	}
	cells := newCellsMatrix(numCols, numRows)
	board := &Board{Cols: numCols, Rows: numRows, Mines: numMines, Cells: cells}
	board = boardInitializer.Initialize(board)
	board.updateNearMines()
	return board, nil
}

// Errors: InvalidBoardProperties
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
func (board Board) updateNearMines() {
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < board.Rows; j++ {
			position := Position{Col: i, Row: j}
			nearMines := board.countNearMines(position)
			cell, _ := board.getCell(position)
			cell.NearMines = nearMines
		}
	}
}

// Complexity: O(1)
// Errors: InvalidPosition
func (board Board) getCell(position Position) (*Cell, error) {
	if position.Col >= 0 && position.Col < board.Cols && position.Row >= 0 && position.Row < board.Rows {
		return board.Cells[position.Col][position.Row], nil
	}
	return nil, modelError.InvalidPosition
}

// Returns up to 8 positions
// Complexity: O(1)
func (board Board) getAdjacencies(position Position) []Position {
	adjacencies := make([]Position, 0, 8)
	for i := position.Col - 1; i <= position.Col+1; i++ {
		for j := position.Row - 1; j <= position.Row+1; j++ {
			// if its a valid column
			if i >= 0 && i <= board.Cols-1 {
				// if its a valid Row
				if j >= 0 && j <= board.Rows-1 {
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

// Returns true if lost game
// Errors: InvalidPosition, ExploreFlagged
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
	if board.Explored == board.Cols*board.Rows-board.Mines {
		board.Solved = true
	}

	// If there are no mines near it, it should explore all adjacent cells
	if currentCell.NearMines == 0 {
		unflaggedNearPositions := board.filterUnflagged(board.getAdjacencies(position))
		for _, unflagged := range unflaggedNearPositions {
			_, _ = board.Explore(unflagged)
		}
	}
	return false, nil
}

// Errors: InvalidPosition
func (board *Board) Flag(position Position, flag Flag) error {
	cell, err := board.getCell(position)
	if err != nil {
		return err
	}
	cell.SetFlag(flag)
	return nil
}

// Complexity: O(1)
func (board Board) filterUnflagged(adjacencies []Position) []Position {
	unflaggedAdjacencies := make([]Position, 0, len(adjacencies))
	for _, adjacency := range adjacencies {
		cell, _ := board.getCell(adjacency)
		if cell.Flag == None {
			unflaggedAdjacencies = append(unflaggedAdjacencies, adjacency)
		}
	}
	return unflaggedAdjacencies
}

// Complexity: O(1)
func (board *Board) countNearMines(position Position) int {
	adjacencies := board.getAdjacencies(position)
	n := 0
	for _, adjacent := range adjacencies {
		cell, _ := board.getCell(adjacent)
		if cell.Mined {
			n++
		}
	}
	return n
}
