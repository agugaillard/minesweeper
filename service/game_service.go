package service

import (
	"github.com/agugaillard/minesweeper/data/cache"
	"github.com/agugaillard/minesweeper/data/redis"
	"github.com/agugaillard/minesweeper/domain/model"
)

type GameService interface {
	NewGame(cols int, rows int, mines int, username model.Username) (*model.Game, error)
	GetGame(id string) (*model.Game, error)
	ExploreCell(gameId string, position model.Position) (*model.Game, error)
	FlagCell(gameId string, position model.Position, flag model.Flag) error
	Save(gameId string) error
	GetOwner(gameId string) (model.Username, error)
}

type DefaultGameService struct {
}

// Errors: InvalidNumberOfMines
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

// Errors: GameNotFound
func (service *DefaultGameService) GetGame(id string) (*model.Game, error) {
	game, err := cache.GameCache.Get(id)
	if err == nil {
		return game, nil
	}
	game, err = redis.GameRedis.Get(id)
	if err != nil {
		return nil, err
	}
	_ = cache.GameCache.New(game)
	return game, nil
}

// Errors: ExploreFlagged, InvalidPosition, GameNotFound, InteractFinished
func (service *DefaultGameService) ExploreCell(gameId string, position model.Position) (*model.Game, error) {
	game, err := service.GetGame(gameId)
	if err != nil {
		return nil, err
	}
	err = game.Explore(position)
	if err != nil {
		return nil, err
	}
	_ = cache.GameCache.Update(game)
	return game, nil
}

// Errors: GameNotFound, InvalidPosition, InteractFinished
func (service *DefaultGameService) FlagCell(gameId string, position model.Position, flag model.Flag) error {
	game, err := service.GetGame(gameId)
	if err != nil {
		return err
	}
	return game.Flag(position, flag)
}

// Errors: GameNotFound
func (service *DefaultGameService) Save(gameId string) error {
	game, err := service.GetGame(gameId)
	if err != nil {
		return err
	}
	err = redis.GameRedis.New(game)
	if err != nil { // if the game already exists
		_ = redis.GameRedis.Update(game)
	}
	return nil
}

func (service *DefaultGameService) GetOwner(gameId string) (model.Username, error) {
	game, err := cache.GameCache.Get(gameId)
	if err != nil {
		return "", err
	}
	return game.Owner, nil
}
