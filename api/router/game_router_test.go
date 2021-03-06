package router

import (
	"bytes"
	"encoding/json"
	"github.com/agugaillard/minesweeper/api/dto"
	"github.com/agugaillard/minesweeper/data/cache"
	"github.com/agugaillard/minesweeper/data/redis"
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
	gameRouter := GameRouter{GameService: &gameServiceTest{}}
	gameRouter.Routes(r)
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.NewGameRequestDto{
		Cols:     3,
		Rows:     3,
		Mines:    2,
		Username: "user",
	})
	req, _ := http.NewRequest("POST", "/game", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}

	var response dto.GameResponseDto
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
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(3, 3, 2, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.ExploreCellRequestDto{
		Position: model.Position{},
		Username: "user",
	})
	req, _ := http.NewRequest("POST", "/game/"+game.Id+"/explore", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}
	var response dto.GameResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error deserializing response")
	}
	if !response.Finished {
		t.Errorf("game should have finished, mine was explored")
	}
	if response.Board.Solved {
		t.Errorf("game shouldn't have been solved, mine was explored")
	}
}

func TestExploreWinGame(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(5, 5, 1, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.ExploreCellRequestDto{
		Position: model.Position{Col: 2, Row: 2},
		Username: "user",
	})
	req, _ := http.NewRequest("POST", "/game/"+game.Id+"/explore", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}
	var response dto.GameResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error deserializing response")
	}
	if !response.Finished {
		t.Errorf("game should have finished, all free cell were explored")
	}
	if !response.Board.Solved {
		t.Errorf("game should have been solved, all free cell were explored")
	}
}

func TestFlagCell(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(5, 5, 1, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.FlagCellRequestDto{
		Position: model.Position{Col: 4, Row: 4},
		Flag:     model.RedFlag,
		Username: "user",
	})
	req, _ := http.NewRequest("PUT", "/game/"+game.Id+"/flag", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}

	if game.Board.Cells[4][4].Flag != model.RedFlag {
		t.Errorf("redflag was expected")
	}
}

func TestExploreFlagged(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(5, 5, 1, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.FlagCellRequestDto{
		Position: model.Position{Col: 0, Row: 0},
		Username: "user",
	})
	req, _ := http.NewRequest("PUT", "/game/"+game.Id+"/flag", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}

	w = httptest.NewRecorder()
	body, _ = json.Marshal(dto.ExploreCellRequestDto{
		Position: model.Position{Col: 0, Row: 0},
		Username: "user",
	})
	req, _ = http.NewRequest("POST", "/game/"+game.Id+"/explore", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 400 was expected")
	}

	var response dto.GameResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error deserializing response")
	}
	if response.Finished {
		t.Errorf("game should have continue, can't explore flagged cell")
	}
}

func TestResume(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(5, 5, 1, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.FlagCellRequestDto{
		Position: model.Position{Col: 4, Row: 4},
		Flag:     model.RedFlag,
		Username: "user",
	})
	req, _ := http.NewRequest("PUT", "/game/"+game.Id+"/flag", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	w = httptest.NewRecorder()
	body, _ = json.Marshal(dto.ResumeGameRequestDto{
		Username: "user",
	})
	req, _ = http.NewRequest("POST", "/game/"+game.Id+"/resume", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("unexpected status code, found " + strconv.Itoa(w.Code) + ", 200 was expected")
	}
	var response dto.GameResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error deserializing response")
	}
	if response.Board.Cells[24].Flag != model.RedFlag {
		t.Errorf("this cell should be flagged with red_flag")
	}
}

func TestExploreAuth(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(5, 5, 1, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.ExploreCellRequestDto{
		Position: model.Position{Col: 4, Row: 4},
		Username: "not_user",
	})
	req, _ := http.NewRequest("POST", "/game/"+game.Id+"/explore", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != 403 {
		t.Errorf("user shouldn't be able to explore in this game")
	}
}

func TestFlagAuth(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(5, 5, 1, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.FlagCellRequestDto{
		Position: model.Position{Col: 4, Row: 4},
		Flag:     model.RedFlag,
		Username: "not_user",
	})
	req, _ := http.NewRequest("PUT", "/game/"+game.Id+"/flag", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != 403 {
		t.Errorf("user shouldn't be able to flag in this game")
	}
}

func TestResumeAuth(t *testing.T) {
	r := gin.Default()
	gameRouter := GameRouter{GameService: GameServiceTest}
	gameRouter.Routes(r)
	game, _ := gameRouter.GameService.New(5, 5, 1, "user")
	w := httptest.NewRecorder()
	body, _ := json.Marshal(dto.ResumeGameRequestDto{
		Username: "not_user",
	})
	req, _ := http.NewRequest("POST", "/game/"+game.Id+"/resume", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != 403 {
		t.Errorf("user shouldn't be able to resume this game")
	}
}

type gameServiceTest struct {
	*service.DefaultGameService
}

var GameServiceTest = &gameServiceTest{DefaultGameService: service.NewDefaultGameService(cache.GameCache, redis.GameRedis)}

func (service *gameServiceTest) New(cols int, rows int, mines int, username model.Username) (*model.Game, error) {
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
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < board.Rows; j++ {
			if minesPlaced < board.Mines {
				board.Cells[i][j] = model.NewMinedCell()
				minesPlaced++
			} else {
				board.Cells[i][j] = model.NewSafeCell()
			}
		}
	}
	return board
}
