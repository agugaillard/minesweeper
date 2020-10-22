package model

type Board struct {
	NumCols  int
	NumRows  int
	NumMines int
	Cells    [][]*Cell
}

func NewBoard(numCols int, numRows int, numMines int, boardInitializer BoardInitializer) *Board {
	cells := newCellsMatrix(numCols, numRows)
	board := &Board{NumCols: numCols, NumRows: numRows, NumMines: numMines, Cells: cells}
	board = boardInitializer.Initialize(board)
	return board
}

func NewRandomBoard(numCols int, numRows int, numMines int) *Board {
	return NewBoard(numCols, numRows, numMines, &RandomMinesBoardInitializer{})
}

func newCellsMatrix(numCols int, numRows int) [][]*Cell {
	cells := make([][]*Cell, numCols)
	for i := range cells {
		cells[i] = make([]*Cell, numRows)
	}
	return cells
}
