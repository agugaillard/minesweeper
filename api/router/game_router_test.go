package router

import (
	"bytes"
	"encoding/json"
	"github.com/agugaillard/minesweeper/api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNewGameRoute(t *testing.T) {
	r := gin.Default()
	GameRoutes(r)
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
}
