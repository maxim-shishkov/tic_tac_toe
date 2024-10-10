package server

import (
	"log"
	"net/http"

	"tic_tac_toe/internal/game/handleres"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr    string
	Handler *handleres.Handler
}

func NewServer(addr string, gameHandler *handleres.Handler) *Server {
	return &Server{
		Addr:    addr,
		Handler: gameHandler,
	}
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/new_game", s.Handler.CreateGame).Methods("POST")
	r.HandleFunc("/move", s.Handler.Move).Methods("POST")

	log.Printf("server is running on %s", s.Addr)
	return http.ListenAndServe(s.Addr, r)
}
