package main

import (
	"flag"

	"github.com/svenfinke/botChallenge/server"
)

var (
	serverPort string
)

func init() {
	flag.StringVar(&serverPort, "p", "8080", "Port")
	flag.Parse()
}

func main() {
	server := server.Server{Port: serverPort}
	server.Run()
}
