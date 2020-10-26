package router

import (
	"bytes"
	"encoding/json"
	"github.com/agugaillard/minesweeper/api/dto"
	"github.com/agugaillard/minesweeper/data/cache"
	"github.com/agugaillard/minesweeper/domain/model"
	"github.com/agugaillard/minesweeper/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNewGameRoute(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: &GameServiceTest{}}
	gameRouter.Routes(r)
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.NewGameRequestDto{
		Cols:  3,
		Rows:  3,
		Mines: 2,
	})
	req, _ := http.NewRequest("POST", "/game", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}

	var response dto.NewGameResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error deserializing response")
	}
	if response.Board.Cols != 3 || response.Board.Rows != 3 || response.Board.Mines != 2 {
		t.Errorf("board properties doesn't match")
	}
}

func TestExploreMine(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: &GameServiceTest{}}
	gameRouter.Routes(r)

	game, _ := gameRouter.GameService.NewGame(3, 3, 2, "")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.ExploreCellRequestDto{
		Position: model.Position{},
	})
	req, _ := http.NewRequest("POST", "/game/"+game.Id+"/explore", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}
	var response dto.NewGameResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error deserializing response")
	}
	if !response.Finished {
		t.Errorf("game should have finished, mine was explored")
	}
}

func TestExploreWinGame(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: &GameServiceTest{}}
	gameRouter.Routes(r)

	game, _ := gameRouter.GameService.NewGame(5, 5, 1, "")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.ExploreCellRequestDto{
		Position: model.Position{Col: 2, Row: 2},
	})
	req, _ := http.NewRequest("POST", "/game/"+game.Id+"/explore", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}
	var response dto.NewGameResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error deserializing response")
	}
	if !response.Finished {
		t.Errorf("game should have finished, all free cell were explored")
	}
}

func TestFlagCell(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: &GameServiceTest{}}
	gameRouter.Routes(r)

	game, _ := gameRouter.GameService.NewGame(5, 5, 1, "")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.FlagCellRequestDto{
		Position: model.Position{Col: 4, Row: 4},
		Flag:     model.RedFlag,
	})
	req, _ := http.NewRequest("POST", "/game/"+game.Id+"/flag", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}

	if game.Board.Cells[4][4].Flag != model.RedFlag {
		t.Errorf("redflag was expected")
	}
}

type GameServiceTest struct {
	service.DefaultGameService
}

func (service *GameServiceTest) NewGame(cols int, rows int, mines int, username model.Username) (*model.Game, error) {
	game, err := model.NewGame(cols, rows, mines, username)
	board, _ := model.NewBoard(cols, rows, mines, &AllTogetherMinesBoardInitializer{})
	game.Board = board
	if err != nil {
		return nil, err
	}
	err = cache.GameCache.New(game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

type AllTogetherMinesBoardInitializer struct{}

func (*AllTogetherMinesBoardInitializer) Initialize(board *model.Board) *model.Board {
	minesPlaced := 0
	for i := 0; i < board.NumCols; i++ {
		for j := 0; j < board.NumRows; j++ {
			if minesPlaced < board.NumMines {
				board.Cells[i][j] = model.NewMinedCell()
				minesPlaced++
			} else {
				board.Cells[i][j] = model.NewSafeCell()
			}
		}
	}
	return board
}
