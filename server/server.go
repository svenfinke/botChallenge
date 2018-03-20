package server

import (
	"log"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type Server struct {
	gameHandler []*GameHandler
	Port        string
	playerIndex int
}

func (server *Server) Run() {
	http.HandleFunc("/", server.handler)
	log.Fatal(http.ListenAndServe(":"+server.Port, nil))
}

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("apikey")
	if !server.validateApiKey(apiKey) {
		w.WriteHeader(401)
		return
	}

	// 1. Create game if none exists
	if server.gameHandler == nil || server.gameHandler[len(server.gameHandler)-1].IsGameReady() {
		id, _ := uuid.NewV4()
		gameHandler := GameHandler{Id: id.String(), Game: &Fourwins{}}
		server.gameHandler = append(server.gameHandler, &gameHandler)
		go gameHandler.Start()
	}

	// 2. Add player to game
	handler := server.gameHandler[len(server.gameHandler)-1]

	// playerId, _ := uuid.NewV1()
	// player := Player{ResponseWriter: &w, Id: playerId.String()}
	server.playerIndex++
	player := Player{ResponseWriter: &w, Id: strconv.Itoa(server.playerIndex), ApiKey: apiKey}

	player.Init()
	handler.Play(&player)
}

func (server *Server) validateApiKey(apiKey string) bool {
	return apiKey != ""
}