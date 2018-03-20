package server

import (
	"fmt"
	"net/http"
)

type Player struct {
	Id             string
	ResponseWriter *http.ResponseWriter
	Request        chan string
	Response       chan string
	Disconnected   bool
	ApiKey         string
}

func (player *Player) Init() {
	http.HandleFunc("/api/"+player.Id, player.handle)
	player.Request = make(chan string)
	player.Response = make(chan string)
}

func (player *Player) handle(w http.ResponseWriter, request *http.Request) {
	if player.Disconnected {
		w.WriteHeader(403)
		return
	}
	apiHeaderKey := request.Header.Get("apikey")
	if apiHeaderKey != player.ApiKey {
		w.WriteHeader(403)
		fmt.Fprint(w, "INVALID APIKEY!")
		player.Disconnected = true
		return
	}
	player.Request <- "Eingang"
	fmt.Fprint(w, <-player.Response)
}
