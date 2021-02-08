package service

import (
	"github.com/agugaillard/minesweeper/domain/model"
	"github.com/agugaillard/minesweeper/domain/repository"
)

type GameService interface {
	New(cols int, rows int, mines int, username model.Username) (*model.Game, error)
	Get(id string) (*model.Game, error)
	ExploreCell(gameId string, position model.Position) (*model.Game, error)
	FlagCell(gameId string, position model.Position, flag model.Flag) error
	GetOwner(gameId string) (model.Username, error)
}

func NewDefaultGameService(cache repository.GameRepository, persistence repository.GameRepository) *DefaultGameService {
	return &DefaultGameService{cache, persistence}
}

type DefaultGameService struct {
	cache       repository.GameRepository
	persistence repository.GameRepository
}

// Errors: InvalidNumberOfMines, GameAlreadyExists
func (service *DefaultGameService) New(cols int, rows int, mines int, username model.Username) (*model.Game, error) {
	game, err := model.NewGame(cols, rows, mines, username)
	if err != nil {
		return nil, err
	}
	err = service.cache.New(game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

// Errors: GameNotFound
func (service *DefaultGameService) Get(id string) (*model.Game, error) {
	game, err := service.cache.Get(id)
	if err == nil {
		return game, nil
	}
	game, err = service.persistence.Get(id)
	if err != nil {
		return nil, err
	}
	_ = service.cache.New(game)
	return game, nil
}

// Errors: ExploreFlagged, InvalidPosition, GameNotFound, InteractFinished
func (service *DefaultGameService) ExploreCell(gameId string, position model.Position) (*model.Game, error) {
	game, err := service.cache.Get(gameId)
	if err != nil {
		return nil, err
	}
	err = game.Explore(position)
	if err != nil {
		return nil, err
	}
	_ = service.cache.Update(game)
	_ = service.persistence.Update(game)
	return game, nil
}

// Errors: GameNotFound, InvalidPosition, InteractFinished
func (service *DefaultGameService) FlagCell(gameId string, position model.Position, flag model.Flag) error {
	game, err := service.Get(gameId)
	if err != nil {
		return err
	}
	err = game.Flag(position, flag)
	if err != nil {
		return err
	}
	_ = service.cache.Update(game)
	_ = service.persistence.Update(game)
	return nil
}

// Errors: GameNotFound
func (service *DefaultGameService) GetOwner(gameId string) (model.Username, error) {
	game, err := service.Get(gameId)
	if err != nil {
		return "", err
	}
	return game.Owner, nil
}
