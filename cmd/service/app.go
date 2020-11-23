package main

import (
	"urlshortener/internal/postgres"
	"urlshortener/internal/transport/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type App struct {
	Store *postgres.Store
	API   *http.Api
	Log   zerolog.Logger
}

func (app *App) Connect() error {
	if err := app.Store.Connect(); err != nil {
		return errors.Wrap(err, "connect to postgres")
	}

	app.Log.Info().Msg("connected to db")

	return nil
}

func (app *App) Close() error {
	if err := app.Store.Close(); err != nil {
		return errors.Wrap(err, "close postgress connection")
	}

	return nil
}
