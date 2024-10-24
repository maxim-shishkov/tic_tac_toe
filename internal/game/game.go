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
	ID       string
	Board    [3][3]string // TODO: для внутреннего представления действительно нужен string?
	Next     string
	Winner   *string
	Finished bool
}

func NewGame(id string) *Game {
	return &Game{
		ID:    id,
		Next:  PlayerX,
		Board: [3][3]string{},
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

	g.Board[row][col] = player

	winner := g.checkWinner()

	if winner != nil && len(*winner) != 0 { // TODO: зачем двойная проверка?
		fmt.Println("winner = ", &winner) // TODO: не взаимодействуем с выводом из бизнес-меотдов
		g.Winner = winner
		g.Finished = true
		return nil
	}

	// TODO: игра так же заканчивается, если не осталось свободных клеток

	g.Next = switchPlayer(player)
	return nil // TODO: удобно получить какой-то полезный результат
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
	first := line[0]
	if first == " " {
		return nil
	}
	for _, cell := range line {
		if cell != first {
			return nil
		}
	}
	winner := first
	return &winner
}
