package server

import (
	"time"
)

type Game interface {
	GetConfig() GameConfig
	Move(player *Player)
	GetGameState(player *Player) string
}

type GameConfig struct {
	PlayerCount int
	Timeout     time.Duration
}
