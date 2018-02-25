package server

import (
	"time"
)

type Fourwins struct {
	id int
}

func (game *Fourwins) GetConfig() GameConfig {
	return GameConfig{2, 1 * time.Second}
}

func (game *Fourwins) Move(player *Player) {
}

func (game *Fourwins) GetGameState(player *Player) string {
	return "Gamestate " + player.Id + "\n"
}