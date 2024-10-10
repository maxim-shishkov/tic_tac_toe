package service

import (
	"github.com/google/uuid"
	"tic_tac_toe/internal/game/domain"
	"tic_tac_toe/internal/game/repository"
)

type Game struct {
	repo *repository.Repository
}

func NewGame(repo *repository.Repository) *Game {
	return &Game{repo: repo}
}

func (g *Game) CreateGame() (*domain.Game, error) {
	game := domain.NewGame(generateID())
	if err := g.repo.Save(game); err != nil {
		return nil, err
	}

	return game, nil
}

func generateID() string {
	return uuid.New().String()
}

func (g *Game) Move(id string, row, col int, player domain.Player) (*domain.Game, error) {
	game, err := g.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if err := game.Move(row, col, player); err != nil {
		return nil, err
	}

	if err := g.repo.Save(game); err != nil {
		return nil, err
	}

	return game, nil
}
