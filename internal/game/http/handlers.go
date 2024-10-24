package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tic_tac_toe/internal/game"
	"tic_tac_toe/internal/server/api"
)

type GameHandler struct {
	gameService *game.Service
}

func NewHandler(service *game.Service) *GameHandler {
	return &GameHandler{gameService: service}
}

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) (any, error) {
	game, err := h.gameService.CreateGame()
	if err != nil {
		return nil, fmt.Errorf("%w:%w", api.ErrInternal, err) // TODO: w должен быть один, так как после этого тип исходной ошибки уже не важен
	}

	return game, nil
}

type MoveRequest struct {
	Id     string `json:"id"`
	Row    int    `json:"row"`
	Col    int    `json:"col"`
	Player string `json:"player"`
}

func (h *GameHandler) Move(w http.ResponseWriter, r *http.Request) (any, error) {
	var req MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("%w:%w", api.ErrBadRequest, err) // TODO
	}

	if req.Row < 0 || req.Row >= 3 {
		return nil, fmt.Errorf("%w:%s", api.ErrBadRequest, "row field is outside")
	}
	if req.Col < 0 || req.Col >= 3 {
		return nil, fmt.Errorf("%w:%s", api.ErrBadRequest, "col field is outside")
	}

	game, err := h.gameService.Move(req.Id, req.Row, req.Col, req.Player)
	if err != nil {
		return nil, err // TODO: обработка ошибок
	}

	return game, nil
}

type StatusRequest struct {
	Id string `json:"id"`
}

func (h *GameHandler) Status(w http.ResponseWriter, r *http.Request) (any, error) {
	var req StatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("%w:%w", api.ErrBadRequest, err) // TODO
	}

	game, err := h.gameService.Status(req.Id)
	if err != nil {
		return nil, fmt.Errorf("%w:%w", api.ErrNotFound, err) // TODO
	}

	return game, nil
}
