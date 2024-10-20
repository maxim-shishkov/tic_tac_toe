package game

import (
	"errors"
)

type Repositories struct {
	games map[string]*Game
}

var ErrNotFound = errors.New("game not found")

func NewRepository() *Repositories {
	return &Repositories{
		games: make(map[string]*Game),
	}
}

func (r *Repositories) Save(game *Game) error {
	r.games[game.ID] = game
	return nil
}

func (r *Repositories) FindByID(id string) (*Game, error) {
	game, exists := r.games[id]
	if !exists {
		return nil, ErrNotFound
	}

	return game, nil
}
