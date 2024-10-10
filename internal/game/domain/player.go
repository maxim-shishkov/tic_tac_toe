package domain

type Player string

const (
	PlayerX Player = "X"
	PlayerO Player = "O"
)

func switchPlayer(player Player) Player {
	if player == PlayerX {
		return PlayerO
	}
	return PlayerX
}
