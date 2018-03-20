package main

import (
	"flag"
	"fmt"

	"github.com/svenfinke/botChallenge/server"
	"github.com/svenfinke/botChallenge/visualizer"
)

var (
	serverPort string
	apiKey     string
	runBot     bool
	visual     bool
)

func init() {
	flag.StringVar(&serverPort, "p", "8080", "Port")
	flag.BoolVar(&runBot, "b", false, "runBot")
	flag.StringVar(&apiKey, "k", "asdasd", "apiKey")
	flag.BoolVar(&visual, "v", false, "visual")
	flag.Parse()
}

func main() {
	if runBot {
		fmt.Println("Running as bot")
		// bot := bot.Bot{ApiKey: apiKey}
		return
	}

	fmt.Println("SERVER started on port: " + serverPort)
	if visual {
		fmt.Println("SERVER starting in visual mode.")
		visualizer.Visualizer()
	}
	server := server.Server{Port: serverPort}
	server.Run()
}
