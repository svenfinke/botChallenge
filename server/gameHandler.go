package server

import (
	"fmt"
	"time"

	"github.com/svenfinke/botChallenge/server/games"
)

type GameHandler struct {
	Id          string
	Game        games.Game
	players     []*Player
	playerIndex int
}

func (handler *GameHandler) addPlayer(player *Player) {
	handler.players = append(handler.players, player)
}

func (handler *GameHandler) IsGameReady() bool {
	return handler.Game.GetConfig().PlayerCount == len(handler.players)
}

func (handler *GameHandler) Play(player *Player) {
	handler.addPlayer(player)
	for !handler.IsGameReady() {
		time.Sleep(200 * time.Millisecond)
	}

	for handler.players[handler.playerIndex] != player {
		time.Sleep(200 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
	fmt.Fprint(*player.ResponseWriter, "Your turn")
	handler.playerIndex++
	handler.playerIndex %= handler.Game.GetConfig().PlayerCount
}
