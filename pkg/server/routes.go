package server

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (server *Server) Routes() Routes {
	return Routes{
		Route{"Index", "GET", "/", server.index},
		Route{"ResultIndex", "POST", "/result", server.resultIndex},
		Route{"ResultShow", "GET", "/result/{id}", server.resultShow},
	}
}
