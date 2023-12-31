package http

import "bphn/artikel-hukum/pkg/server"

type API struct {
	server *server.HttpServer
}

func NewAPI(httpServer *server.HttpServer) *API {
	return &API{server: httpServer}
}
