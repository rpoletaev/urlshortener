package cmd

import (
	"urlshortener/internal/postgres"
	"urlshortener/internal/transport/http"

	"github.com/rs/zerolog"
)

type App struct {
	Store *postgres.Store
	API   *http.Api
	Log   zerolog.Logger
}
