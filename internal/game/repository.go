package game

import (
	"errors"
)

type Repository struct {
	games map[string]*Game
}

var ErrNotFound = errors.New("game not found")

func NewRepository() *Repository {
	return &Repository{
		games: make(map[string]*Game),
	}
}

func (r *Repository) Save(game *Game) error {
	r.games[game.ID] = game
	return nil
}

func (r *Repository) FindByID(id string) (*Game, error) {
	game, exists := r.games[id]
	if !exists {
		return nil, ErrNotFound
	}

	return game, nil
}
