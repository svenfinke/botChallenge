package games

type Game interface {
	GetConfig() GameConfig
	Move()
}

type GameConfig struct {
	PlayerCount int
}
