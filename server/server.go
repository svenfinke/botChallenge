package server

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/svenfinke/botChallenge/server/games"
)

type Server struct {
	gameHandler []*GameHandler
}

func (server *Server) Run() {
	http.HandleFunc("/", server.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	// 1. Create game if none exists
	if server.gameHandler == nil || server.gameHandler[len(server.gameHandler)-1].IsGameReady() {
		id, _ := uuid.NewV4()
		server.gameHandler = append(server.gameHandler, &GameHandler{Id: id.String(), Game: &games.Fourwins{}})
	}

	// 2. Add player to game
	handler := server.gameHandler[len(server.gameHandler)-1]
	handler.Play(&Player{ResponseWriter: &w})
}
