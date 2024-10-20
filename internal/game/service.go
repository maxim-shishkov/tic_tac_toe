package game

import (
	"github.com/google/uuid"
)

type Service struct {
	repo *Repositories
}

func NewService(repo *Repositories) *Service {
	return &Service{repo: repo}
}

func (g *Service) CreateGame() (*Game, error) {
	game := NewGame(generateID())
	if err := g.repo.Save(game); err != nil {
		return nil, err
	}

	return game, nil
}

func generateID() string {
	return uuid.New().String()
}

func (g *Service) Move(id string, row, col int, player string) (*Game, error) {
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

func (g *Service) Status(id string) (*Game, error) {
	game, err := g.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return game, nil
}
