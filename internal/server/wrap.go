package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"tic_tac_toe/internal/game"
	"tic_tac_toe/internal/server/api"
)

func wrap(handler func(http.ResponseWriter, *http.Request) (any, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		result, err := handler(w, r)
		if err != nil {
			handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func handleError(w http.ResponseWriter, err error) {
	httpStatusCode := http.StatusInternalServerError
	switch true {
	case errors.Is(err, api.ErrNotFound):
		httpStatusCode = http.StatusNotFound
	case errors.Is(err, api.ErrBadRequest) || errors.Is(err, game.ErrFinished) ||
		errors.Is(err, game.ErrOccupied) || errors.Is(err, game.ErrNotPlayer):
		httpStatusCode = http.StatusBadRequest

	case errors.Is(err, api.ErrInternal):
		httpStatusCode = http.StatusInternalServerError
	}

	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(api.Error{
		Code: httpStatusCode,
		Err:  err.Error(),
	})
}
