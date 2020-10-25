package service

import (
	"github.com/agugaillard/minesweeper/data/cache"
	"github.com/agugaillard/minesweeper/domain/model"
)

type GameService interface {
	NewGame(cols int, rows int, mines int, username model.Username) (*model.Game, error)
	GetGame(id string) (*model.Game, error)
	ExploreCell(gameId string, position model.Position) (*model.Game, error)
}

type DefaultGameService struct {
}

func (service *DefaultGameService) NewGame(cols int, rows int, mines int, username model.Username) (*model.Game, error) {
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

func (service *DefaultGameService) GetGame(id string) (*model.Game, error) {
	game, err := cache.GameCache.Get(id)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (service *DefaultGameService) ExploreCell(gameId string, position model.Position) (*model.Game, error) {
	game, err := service.GetGame(gameId)
	if err != nil {
		return nil, err
	}
	_ = game.Explore(position)
	_ = cache.GameCache.Update(game)
	return game, nil
}
