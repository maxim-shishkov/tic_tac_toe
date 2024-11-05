package game

const (
	PlayerX string = "X"
	PlayerO string = "O"
)

func (g *Game) switchPlayer(player string) {
	if player == g.SymbolX {
		g.Next = g.SymbolO
	}
	g.Next = g.SymbolX
}
