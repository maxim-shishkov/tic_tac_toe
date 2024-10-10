package repository

import (
	"errors"

	"tic_tac_toe/internal/game/domain"
)

type Repository struct {
	games map[string]*domain.Game
}

func NewRepository() *Repository {
	return &Repository{
		games: make(map[string]*domain.Game),
	}
}

func (r *Repository) Save(game *domain.Game) error {
	r.games[game.ID] = game
	return nil
}

func (r *Repository) FindByID(id string) (*domain.Game, error) {
	game, exists := r.games[id]
	if !exists {
		return nil, errors.New("game not found")
	}

	return game, nil
}
