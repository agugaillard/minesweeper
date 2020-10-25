package cache

import (
	"errors"
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

func (cache *gameCache) New(game *model.Game) error {
	if _, err := cache.Get(game.Id); err == nil {
		return errors.New("game already exists")
	}
	cache.games[game.Id] = game
	return nil
}

func (cache *gameCache) Get(id string) (*model.Game, error) {
	game, ok := cache.games[id]
	if !ok {
		return nil, errors.New("game not found")
	}
	return game, nil
}

func (cache *gameCache) Remove(id string) error {
	if _, err := cache.Get(id); err != nil {
		return err
	}
	delete(cache.games, id)
	return nil
}

func (cache *gameCache) GetOwner(id string) (model.Username, error) {
	game, err := cache.Get(id)
	if err != nil {
		return "", err
	}
	return game.Owner, nil
}

func (cache *gameCache) Update(game *model.Game) error {
	if _, err := cache.Get(game.Id); err != nil {
		return err
	}
	cache.games[game.Id] = game
	return nil
}
