package cache

import (
	dataError "github.com/agugaillard/minesweeper/data/error"
	"github.com/agugaillard/minesweeper/domain/model"
)

type gameCache struct {
	games map[string]*model.Game
}

var GameCache = newGameCache()

func newGameCache() *gameCache {
	m := make(map[string]*model.Game, 0)
	return &gameCache{m}
}

// Errors: GameAlreadyExists
func (cache *gameCache) New(game *model.Game) error {
	if _, err := cache.Get(game.Id); err == nil {
		return dataError.GameAlreadyExists
	}
	cache.games[game.Id] = game
	return nil
}

// Errors: GameNotFound
func (cache *gameCache) Get(id string) (*model.Game, error) {
	game, ok := cache.games[id]
	if !ok {
		return nil, dataError.GameNotFound
	}
	return game, nil
}

// Errors: GameNotFound
func (cache *gameCache) Remove(id string) error {
	if _, err := cache.Get(id); err != nil {
		return err
	}
	delete(cache.games, id)
	return nil
}

// Errors: GameNotFound
func (cache *gameCache) GetOwner(id string) (model.Username, error) {
	game, err := cache.Get(id)
	if err != nil {
		return "", err
	}
	return game.Owner, nil
}

// Errors: GameNotFound
func (cache *gameCache) Update(game *model.Game) error {
	if _, err := cache.Get(game.Id); err != nil {
		return err
	}
	cache.games[game.Id] = game
	return nil
}
