package http

import (
	"net/http"
	"urlshortener/internal/service"

	"github.com/gorilla/mux"
)

type Config struct {
	Port string `envconfig:"PORT"`
}

type IPExtractor interface {
	GetIP(r *http.Request) string
}

type RealIPExtractor struct{}

const forwardedHeader = "X-FORWARDED-FOR"

func (ext RealIPExtractor) GetIP(r *http.Request) string {
	forwarded := r.Header.Get(forwardedHeader)
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

type Api struct {
	*Config
	Svc         *service.Service
	IpExtractor IPExtractor
	server      *http.Server `wire:"-"`
}

func (api *Api) Server() *http.Server {
	if api.server != nil {
		return api.server
	}

	r := mux.NewRouter()

	linkrouter := r.PathPrefix("/link").Subrouter()
	linkrouter.HandleFunc("/{hash}", api.getLink).Methods(http.MethodGet)
	linkrouter.HandleFunc("/", api.addLink).Methods(http.MethodPost)

	statrouter := r.PathPrefix("/stat").Subrouter()
	statrouter.HandleFunc("/ip", api.getIPStat).Queries("from", "{from}").Queries("to", "{to}").Methods(http.MethodGet)
	statrouter.HandleFunc("/url", api.getURLStat).Queries("from", "{from}").Queries("to", "{to}").Methods(http.MethodGet)

	api.server = &http.Server{
		Addr:    api.Port,
		Handler: r, // TODO: needs to set up a middlewares
	}

	return api.server
}
