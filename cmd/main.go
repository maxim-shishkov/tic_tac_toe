package main

import (
	"flag"
	"fmt"
	"log"

	"tic_tac_toe/internal/game"
	"tic_tac_toe/internal/game/http"
	"tic_tac_toe/internal/server"
)

var port *int

func init() {
	port = flag.Int("port", 8080, "server port")
	flag.Parse()
}

func main() {
	repo := game.NewRepository()
	service := game.NewService(repo)
	gameHandler := http.NewHandler(service)

	srv := server.NewServer(fmt.Sprintf(":%d", *port), gameHandler)

	if err := srv.Run(); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
