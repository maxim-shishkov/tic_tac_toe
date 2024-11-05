package game

import (
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (g *Service) CreateGame(symbolX, symbolO *string) (*Game, error) {
	game := NewGame(generateID())

	if symbolX != nil {
		game.Next = *symbolX
		game.SymbolX = *symbolX
	}

	if symbolO != nil {
		game.SymbolO = *symbolO
	}

	if err := g.repo.Save(game); err != nil {
		return nil, fmt.Errorf("repo.Save: %w", err)
	}

	return game, nil
}

func generateID() string {
	return uuid.New().String()
}

func (g *Service) Move(id string, row, col int, player string) (*Game, error) {
	game, err := g.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("repo.FindByID: %w", err)
	}

	if err := game.Move(row, col, player); err != nil {
		return nil, fmt.Errorf("repo.Move: %w", err)
	}

	if err := g.repo.Save(game); err != nil {
		return nil, fmt.Errorf("repo.Save: %w", err)
	}

	return game, nil
}

func (g *Service) Status(id string) (*Game, error) {
	game, err := g.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("repo.FindByID: %w", err)
	}

	return game, nil
}
