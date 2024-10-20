package game

const (
	PlayerX string = "X"
	PlayerO string = "O"
)

func switchPlayer(player string) string {
	if player == PlayerX {
		return PlayerO
	}
	return PlayerX
}
