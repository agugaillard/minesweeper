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
	minesPositions := rand.Perm(board.Rows * board.Cols)[:board.Mines]
	isMine := make(map[int]bool, board.Mines)
	for _, minePosition := range minesPositions {
		isMine[minePosition] = true
	}
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < board.Rows; j++ {
			position := i*board.Rows + j
			if isMine[position] {
				board.Cells[i][j] = NewMinedCell()
			} else {
				board.Cells[i][j] = NewSafeCell()
			}
		}
	}
	return board
}
