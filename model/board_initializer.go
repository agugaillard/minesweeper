package model

import (
	"math/rand"
	"time"
)

type BoardInitializer interface {
	Initialize(*Board) *Board
}

type RandomMinesBoardInitializer struct {
}

// Complexity: O(m * n)
func (*RandomMinesBoardInitializer) Initialize(board *Board) *Board {
	rand.Seed(time.Now().UnixNano())
	minesPositions := rand.Perm(board.NumRows * board.NumCols)[:board.NumMines]
	isMine := make(map[int]bool, board.NumMines)
	for _, minePosition := range minesPositions {
		isMine[minePosition] = true
	}
	for i := 0; i < board.NumCols; i++ {
		for j := 0; j < board.NumRows; j++ {
			position := i*board.NumRows + j
			if isMine[position] {
				board.Cells[i][j] = NewMinedCell()
			} else {
				board.Cells[i][j] = NewSafeCell()
			}
		}
	}
	return board
}
