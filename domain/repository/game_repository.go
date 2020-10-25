package repository

import "github.com/agugaillard/minesweeper/domain/model"

type GameRepository interface {
	New(*model.Game) error
	Get(id string) (*model.Game, error)
	Remove(id string) error
	GetOwner(id string) (model.Username, error)
	Update(*model.Game) error
}

var gameRepository GameRepository

func GetGameRepository() GameRepository {
	return gameRepository
}

func InitGameRepository(gr GameRepository) {
	gameRepository = gr
}
