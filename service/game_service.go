package service

import (
	"github.com/agugaillard/minesweeper/data/cache"
	"github.com/agugaillard/minesweeper/domain/model"
)

func NewGame(cols int, rows int, mines int, username model.Username) (*model.Game, error) {
	game, err := model.NewGame(cols, rows, mines, username)
	if err != nil {
		return nil, err
	}
	err = cache.GameCache.New(game)
	if err != nil {
		return nil, err
	}
	return game, nil
}
