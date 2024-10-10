package domain

import (
	"errors"
)

type Game struct {
	ID       string
	Board    [3][3]string
	Next     Player
	Winner   *Player
	Finished bool
}

func NewGame(id string) *Game {
	return &Game{
		ID:    id,
		Next:  PlayerX,
		Board: [3][3]string{},
	}
}

func (g *Game) Move(row, col int, player Player) error {
	if g.Finished {
		return errors.New("game is finished")
	}
	if g.Board[row][col] != "" {
		return errors.New("cell and row is occupied")
	}
	if g.Next != player {
		return errors.New("not your player")
	}

	g.Board[row][col] = string(player)

	winner := g.checkWinner()
	if winner != nil {
		g.Winner = winner
		g.Finished = true
		return nil
	}

	g.Next = switchPlayer(player)
	return nil
}

func (g *Game) checkWinner() *Player {
	// todo
	return nil
}
