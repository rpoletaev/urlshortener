package http

import "urlshortener/internal/service"

type Config struct {
	Port string `envconfig:"PORT"`
}

type Api struct {
	*Config
	Svc *service.Service
}
