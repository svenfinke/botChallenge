package server

import "net/http"

type Player struct {
	id             string
	ResponseWriter *http.ResponseWriter
}
