package server

import "net/http"

type Route struct {
	Name    string
	Methods string
	Path    string
	Handler http.HandlerFunc
}

func (server *Server) Routes() []Route {
	return []Route{
		{"Index", "GET", "/", server.index},
		{"ResultIndex", "POST", "/resource", server.resourceIndex},
		{"ResultShow", "GET", "/resource/{id}", server.resourceShow},
	}
}
