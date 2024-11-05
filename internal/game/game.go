package game

import (
	"errors"
	"fmt"
)

var (
	ErrFinished  = errors.New("game is finished")
	ErrOccupied  = errors.New("cell and row is occupied")
	ErrNotPlayer = errors.New("not your player")
)

type Game struct {
	ID       string       `json:"id"`
	Board    [3][3]string `json:"board"`
	Next     string       `json:"next"`
	Winner   *string      `json:"winner"`
	Finished bool         `json:"finished"`
	SymbolX  string       `json:"symbol_X"`
	SymbolO  string       `json:"symbol_O"`
}

func NewGame(id string) *Game {
	return &Game{
		ID:      id,
		Next:    PlayerX,
		Board:   [3][3]string{},
		SymbolX: PlayerX,
		SymbolO: PlayerO,
	}
}

func (g *Game) Move(row, col int, player string) error {
	if g.Finished {
		return ErrFinished
	}
	if g.Board[row][col] != "" {
		return ErrOccupied
	}
	if g.Next != player {
		return ErrNotPlayer
	}

	g.Board[row][col] = string(player)

	winner := g.checkWinner()

	if winner != nil {
		g.Winner = winner
		g.Finished = true
		return nil
	}
	if g.noMoreSpaces() {
		return fmt.Errorf("free space on the field is over")
	}

	g.switchPlayer(player)

	return nil
}

func (g *Game) noMoreSpaces() bool {
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			if len(g.Board[i][j]) == 0 {
				return false
			}
		}
	}
	return true
}

func (g *Game) checkWinner() *string {
	n := len(g.Board)

	for i := 0; i < n; i++ {
		row := make([]string, n)
		col := make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = g.Board[i][j]
			col[j] = g.Board[j][i]
		}
		if winner := checkWinningLine(row); winner != nil {
			return winner
		}
		if winner := checkWinningLine(col); winner != nil {
			return winner
		}
	}

	diagL := make([]string, n)
	diagR := make([]string, n)
	for i := 0; i < n; i++ {
		diagL[i] = g.Board[i][i]
		diagR[i] = g.Board[i][n-i-1]
	}
	if winner := checkWinningLine(diagL); winner != nil {
		return winner
	}
	if winner := checkWinningLine(diagR); winner != nil {
		return winner
	}

	return nil
}

func checkWinningLine(line []string) *string {
	if len(line[0]) != 0 && line[0] == line[1] && line[1] == line[2] {
		return &line[0]
	}
	return nil
}
