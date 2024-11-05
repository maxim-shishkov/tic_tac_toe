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
	var apiErr *api.Error
	httpStatusCode := http.StatusInternalServerError

	switch {
	case errors.Is(err,
		game.ErrFinished),
		errors.Is(err, game.ErrOccupied),
		errors.Is(err, game.ErrNotPlayer):
		apiErr = api.BadRequest("bad request", err)
		httpStatusCode = http.StatusBadRequest

	default:
		if customErr, ok := err.(*api.Error); ok {
			apiErr = customErr
			switch customErr.Code {
			case api.CodeNotFound:
				httpStatusCode = http.StatusNotFound
			case api.CodeBadRequest:
				httpStatusCode = http.StatusBadRequest
			case api.CodeInternal:
				httpStatusCode = http.StatusInternalServerError
			default:
				httpStatusCode = http.StatusInternalServerError
			}
		} else {
			apiErr = api.InternalError("internal server error", err)
			httpStatusCode = http.StatusInternalServerError
		}
	}

	w.WriteHeader(httpStatusCode)
	if encodeErr := json.NewEncoder(w).Encode(apiErr); encodeErr != nil {
		http.Error(w, "failed to encode error response", http.StatusInternalServerError)
	}
}
