package server

import (
	"fmt"
	"time"
)

type GameHandler struct {
	Id          string
	Game        Game
	players     []*Player
	playerIndex int
	turn        int
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
	// Send to first player
	if handler.players[0] == player {
		fmt.Fprintf(*player.ResponseWriter, handler.Game.GetGameState(player))
	}
}

func (handler *GameHandler) Start() {
	for !handler.IsGameReady() {
		time.Sleep(200 * time.Millisecond)
	}
	handler.Loop()
}

func (handler *GameHandler) Loop() {
	for {
		actualPlayer := handler.players[handler.playerIndex]
		nextIndex := (handler.playerIndex + 1) % handler.Game.GetConfig().PlayerCount
		nextPlayer := handler.players[nextIndex]
		select {
		case <-actualPlayer.Request:
			// Players first move
			handler.Game.Move(actualPlayer)
			if handler.turn < nextIndex {
				fmt.Fprintf(*nextPlayer.ResponseWriter, handler.Game.GetGameState(nextPlayer))
			} else {
				nextPlayer.Response <- handler.Game.GetGameState(nextPlayer)
			}
		case <-time.After(handler.Game.GetConfig().Timeout):
			fmt.Printf("Player (%v) timedout\n", actualPlayer.Id)
			if handler.turn < nextIndex {
				fmt.Fprintf(*nextPlayer.ResponseWriter, handler.Game.GetGameState(nextPlayer))
			} else {
				nextPlayer.Response <- handler.Game.GetGameState(nextPlayer)
			}
			
		}
		handler.playerIndex = nextIndex
		handler.turn++
	}
}
