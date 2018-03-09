package main

import (
	"flag"
	"fmt"

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
	fmt.Println("SERVER started on port: " + serverPort)
	server := server.Server{Port: serverPort}
	server.Run()
}
