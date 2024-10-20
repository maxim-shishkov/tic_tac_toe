package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	gHttp "tic_tac_toe/internal/game/http"
)

type Server struct {
	addr        string
	gameHandler *gHttp.GameHandler
}

func NewServer(addr string, gameHandler *gHttp.GameHandler) *Server {
	return &Server{
		addr:        addr,
		gameHandler: gameHandler,
	}
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/new_game", wrap(s.gameHandler.CreateGame)).Methods("POST")
	r.HandleFunc("/move", wrap(s.gameHandler.Move)).Methods("POST")
	r.HandleFunc("/status", wrap(s.gameHandler.Status)).Methods("POST")

	log.Printf("server is running on %s", s.addr)
	return http.ListenAndServe(s.addr, r)
}
