package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gHttp "tic_tac_toe/internal/game/http"

	"github.com/gorilla/mux"
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

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", s.addr),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Printf("server is running on %s", s.addr)
		if err := http.ListenAndServe(s.addr, r); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("srv.ListenAndServe: %s", err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 2*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		srv.Close()

		return fmt.Errorf("srv.Shutdown: %w", err)
	}

	log.Println("server stop.")
	return nil
}
