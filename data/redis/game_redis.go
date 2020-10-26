package redis

import (
	"context"
	"encoding/json"
	dataError "github.com/agugaillard/minesweeper/data/error"
	"github.com/agugaillard/minesweeper/domain/model"
	"github.com/go-redis/redis/v8"
)

type gameRedis struct {
	rdb *redis.Client
}

var (
	ctx       = context.Background()
	GameRedis = &gameRedis{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
)

// Errors: GameAlreadyExists
func (gr *gameRedis) New(game *model.Game) error {
	_, err := gr.Get(game.Id)
	if err == nil {
		return dataError.GameAlreadyExists
	}
	// Warning: assume the data is always valid
	gameStr, _ := json.Marshal(game)
	gr.rdb.Set(ctx, game.Id, gameStr, 0)
	return nil
}

// Errors: GameNotFound
func (gr *gameRedis) Get(id string) (*model.Game, error) {
	gameStr, err := gr.rdb.Get(ctx, id).Result()
	if err != nil {
		return nil, dataError.GameNotFound
	}
	var game *model.Game
	// Warning: Assume the data is always valid
	_ = json.Unmarshal([]byte(gameStr), &game)
	return game, nil
}

// Errors: GameNotFound
func (gr *gameRedis) Remove(id string) error {
	if _, err := gr.Get(id); err != nil {
		return err
	}
	gr.rdb.Del(ctx, id)
	return nil
}

// Errors: GameNotFound
func (gr *gameRedis) GetOwner(id string) (model.Username, error) {
	game, err := gr.Get(id)
	if err != nil {
		return "", err
	}
	return game.Owner, nil
}

// Errors: GameNotFound
func (gr *gameRedis) Update(game *model.Game) error {
	if _, err := gr.Get(game.Id); err != nil {
		return err
	}
	// Warning: assume the data is always valid
	gameStr, _ := json.Marshal(game)
	gr.rdb.Set(ctx, game.Id, gameStr, 0)
	return nil
}
