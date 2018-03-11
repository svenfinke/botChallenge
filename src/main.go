package main

import (
	"flag"
	"fmt"

	"github.com/svenfinke/botChallenge/bot"
	"github.com/svenfinke/botChallenge/server"
)

var (
	serverPort string
	runBot     bool
	apiKey     string
)

func init() {
	flag.StringVar(&serverPort, "p", "8080", "Port")
	flag.BoolVar(&runBot, "b", false, "runBot")
	flag.StringVar(&apiKey, "k", "asdasd", "apiKey")
	flag.Parse()
}

func main() {
	if runBot {
		fmt.Println("Running as bot")
		bot = bot.Bot{ApiKey: apiKey}
		return
	}

	fmt.Println("SERVER started on port: " + serverPort)
	server := server.Server{Port: serverPort}
	server.Run()
}
