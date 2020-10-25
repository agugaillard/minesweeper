package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/agugaillard/minesweeper/domain/model"
	"github.com/go-redis/redis/v8"
)

type gameRedis struct {
}

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx       = context.Background()
	GameRedis = gameRedis{}
)

func (gr *gameRedis) New(game *model.Game) error {
	_, err := gr.Get(game.Id)
	if err == nil {
		return errors.New("game already exists")
	}
	// Warning: assume the data is always valid
	gameStr, _ := json.Marshal(game)
	rdb.Set(ctx, game.Id, gameStr, 0)
	return nil
}

func (gr *gameRedis) Get(id string) (*model.Game, error) {
	gameStr, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return nil, errors.New("game not found")
	}
	var game *model.Game
	// Warning: Assume the data is always valid
	_ = json.Unmarshal([]byte(gameStr), &game)
	return game, nil
}

func (gr gameRedis) Remove(id string) error {
	if _, err := gr.Get(id); err != nil {
		return err
	}
	rdb.Del(ctx, id)
	return nil
}

func (gr gameRedis) GetOwner(id string) (model.Username, error) {
	game, err := gr.Get(id)
	if err != nil {
		return "", err
	}
	return game.Owner, nil
}

func (gr gameRedis) Update(game *model.Game) error {
	if _, err := gr.Get(game.Id); err != nil {
		return err
	}
	// Warning: assume the data is always valid
	gameStr, _ := json.Marshal(game)
	rdb.Set(ctx, game.Id, gameStr, 0)
	return nil
}
