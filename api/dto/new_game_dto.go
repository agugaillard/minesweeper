package dto

import (
	"github.com/agugaillard/minesweeper/domain/model"
	"time"
)

type NewGameRequestDto struct {
	Cols     int            `json:"cols"`
	Rows     int            `json:"rows"`
	Mines    int            `json:"mines"`
	Username model.Username `json:"username"`
}

type GameResponseDto struct {
	Id       string            `json:"id"`
	Start    time.Time         `json:"start"`
	End      time.Time         `json:"end"`
	Owner    model.Username    `json:"owner"`
	Board    *boardResponseDto `json:"board"`
	Finished bool              `json:"finished"`
}

func NewGameDto(game *model.Game) *GameResponseDto {
	return &GameResponseDto{
		Id:       game.Id,
		Start:    game.Start,
		End:      game.End,
		Owner:    game.Owner,
		Board:    newBoardDto(game.Board),
		Finished: game.Finished,
	}
}

type boardResponseDto struct {
	Cols     int                `json:"cols"`
	Rows     int                `json:"rows"`
	Mines    int                `json:"mines"`
	Cells    []*CellResponseDto `json:"cells"`
	Explored int                `json:"explored"`
	Solved   bool               `json:"solved"`
}

func newBoardDto(board *model.Board) *boardResponseDto {
	return &boardResponseDto{
		Cols:     board.Cols,
		Rows:     board.Rows,
		Mines:    board.Mines,
		Cells:    matrixToArray(board.Cells),
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

func matrixToArray(cells [][]*model.Cell) []*CellResponseDto {
	arrayDto := make([]*CellResponseDto, 0, len(cells)*len(cells[0]))
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[0]); j++ {
			cellDto := NewCellDto(cells[i][j])
			cellDto.Position = model.Position{Col: i, Row: j}
			arrayDto = append(arrayDto, cellDto)
		}
	}
	return arrayDto
}
