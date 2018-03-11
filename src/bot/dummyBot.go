package bot

import (
	"fmt"
	"math/rand"
	"time"
)

type Bot struct {
	ApiKey string
	moves  []string
}

func (bot *Bot) Init() {
	rand.Seed(time.Now().Unix())
	bot.moves = append(bot.moves, "1")
	bot.moves = append(bot.moves, "2")
	bot.moves = append(bot.moves, "3")
	bot.moves = append(bot.moves, "4")
}

func (bot *Bot) RunBot() {
	fmt.Println("Bot started with key: " + bot.apiKey)
}

func (bot *Bot) Move() string {
	return bot.moves[rand.Intn(len(bot.moves))]
}
