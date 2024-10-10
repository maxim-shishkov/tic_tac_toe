package handleres

import (
	"encoding/json"
	"net/http"

	"tic_tac_toe/internal/game/domain"
	"tic_tac_toe/internal/game/service"
)

type Handler struct {
	game *service.Game
}

func NewHandler(g *service.Game) *Handler {
	return &Handler{game: g}
}

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	game, err := h.game.CreateGame()
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(game)
}

type MoveRequest struct {
	Id     string        `json:"id"`
	Row    int           `json:"row"`
	Col    int           `json:"col"`
	Player domain.Player `json:"player"`
}

func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	var req MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	if req.Row < 0 || req.Row >= 3 {
		return
	}
	if req.Col < 0 || req.Col >= 3 {
		return
	}

	game, err := h.game.Move(req.Id, req.Row, req.Col, req.Player)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(game)
}
