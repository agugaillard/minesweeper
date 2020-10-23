package dto

import (
	"github.com/agugaillard/minesweeper/domain/model"
	"time"
)

type NewGameRequestDto struct {
	Cols  int `json:"cols"`
	Rows  int `json:"rows"`
	Mines int `json:"mines"`
}

type NewGameResponseDto struct {
	Id       string            `json:"id"`
	Start    time.Time         `json:"start"`
	End      time.Time         `json:"end,omitempty"`
	Owner    model.Username    `json:"owner"`
	Board    *boardResponseDto `json:"board"`
	Finished bool              `json:"finished"`
}

func NewGameDto(game *model.Game) *NewGameResponseDto {
	return &NewGameResponseDto{
		Id:       game.Id,
		Start:    game.Start,
		End:      game.End,
		Owner:    game.Owner,
		Board:    newBoardDto(game.Board),
		Finished: game.Finished,
	}
}

type boardResponseDto struct {
	Cols     int                  `json:"cols"`
	Rows     int                  `json:"rows"`
	Mines    int                  `json:"mines"`
	Cells    [][]*CellResponseDto `json:"cells"`
	Explored int                  `json:"explored"`
	Solved   bool                 `json:"solved"`
}

func newBoardDto(board *model.Board) *boardResponseDto {
	return &boardResponseDto{
		Cols:     board.NumCols,
		Rows:     board.NumRows,
		Mines:    board.NumMines,
		Cells:    newCellsDtoMatrix(board.Cells),
		Explored: board.Explored,
		Solved:   board.Solved,
	}
}

type CellResponseDto struct {
	model.Position
	Explored  bool       `json:"explored"`
	Flag      model.Flag `json:"flag"`
	NearMines int        `json:"near_mines"`
}

func NewCellDto(cell *model.Cell) *CellResponseDto {
	nearMines := -1
	if cell.Explored {
		nearMines = cell.NearMines
	}
	return &CellResponseDto{
		Explored:  cell.Explored,
		Flag:      cell.Flag,
		NearMines: nearMines,
	}
}

func newCellsDtoMatrix(cells [][]*model.Cell) [][]*CellResponseDto {
	cols := len(cells)
	rows := len(cells[0])
	cellsDto := make([][]*CellResponseDto, cols)
	for i := range cells {
		cellsDto[i] = make([]*CellResponseDto, rows)
	}
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			cellsDto[i][j] = NewCellDto(cells[i][j])
			cellsDto[i][j].Col, cellsDto[i][j].Col = i, j
		}
	}
	return cellsDto
}
