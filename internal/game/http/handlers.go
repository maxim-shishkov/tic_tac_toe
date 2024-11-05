package http

import (
	"encoding/json"
	"errors"
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

type CreateRequest struct {
	SymbolX *string `json:"symbol_x"`
	SymbolO *string `json:"symbol_o"`
}

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) (any, error) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, api.BadRequest("json.Decode", err)
	}

	game, err := h.gameService.CreateGame(req.SymbolX, req.SymbolO)
	if err != nil {
		return nil, api.InternalError("createGame", err)
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
		return nil, api.BadRequest("json.Decode", err)
	}

	if req.Row < 0 || req.Row >= 3 {
		return nil, api.BadRequest("req.Row", errors.New("row field is outside"))
	}

	if req.Col < 0 || req.Col >= 3 {
		return nil, api.BadRequest("req.Col", errors.New("col field is outside"))
	}

	g, err := h.gameService.Move(req.Id, req.Row, req.Col, req.Player)
	if err != nil {
		if errors.Is(err, game.ErrNotFound) {
			return nil, api.NotFound("game not found")
		}
		return nil, api.InternalError("gameService.Move", err)
	}

	return g, nil
}

type StatusRequest struct {
	Id string `json:"id"`
}

func (h *GameHandler) Status(w http.ResponseWriter, r *http.Request) (any, error) {
	var req StatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, api.BadRequest("json.Decode", err)
	}

	g, err := h.gameService.Status(req.Id)
	if err != nil {
		if errors.Is(err, game.ErrNotFound) {
			return nil, api.NotFound("game not found")
		}
		return nil, api.InternalError("gameService.Status", err)
	}

	return g, nil
}
