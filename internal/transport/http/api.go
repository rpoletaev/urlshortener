package http

import (
	"net/http"
	"urlshortener/internal/service"

	"github.com/gorilla/mux"
)

type Config struct {
	Port string `envconfig:"PORT"`
}

type Api struct {
	*Config
	Svc    *service.Service
	server *http.Server `wire:"-"`
}

func (api *Api) Server() *http.Server {
	if api.server != nil {
		return api.server
	}

	r := mux.NewRouter()

	linkrouter := r.PathPrefix("/link").Subrouter()
	linkrouter.HandleFunc("/{hash}", api.getLink).Methods(http.MethodGet)
	linkrouter.HandleFunc("/", api.addLink).Methods(http.MethodPost)

	api.server = &http.Server{
		Addr:    api.Port,
		Handler: r, // TODO: needs to set up a middlewares
	}

	return api.server
}
