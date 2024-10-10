package main

import (
	"fmt"
	"log"

	"tic_tac_toe/internal/game/handleres"
	"tic_tac_toe/internal/game/repository"
	"tic_tac_toe/internal/game/service"
	"tic_tac_toe/server"
)

func main() {
	repo := repository.NewRepository()
	game := service.NewGame(repo)
	gameHandler := handleres.NewHandler(game)

	port := 8080 // todo
	srv := server.NewServer(fmt.Sprintf(":%d", port), gameHandler)

	if err := srv.Run(); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
