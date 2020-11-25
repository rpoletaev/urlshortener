package main

import (
	"urlshortener/internal/postgres"
	"urlshortener/internal/redis"
	"urlshortener/internal/transport/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type App struct {
	Store *postgres.Store
	API   *http.Api
	Redis *redis.Backend
	Log   zerolog.Logger
}

func (app *App) Connect() error {
	if err := app.Store.Connect(); err != nil {
		return errors.Wrap(err, "connect to postgres")
	}

	app.Log.Info().Msg("connected to db")

	if err := app.Redis.Connect(); err != nil {
		return errors.Wrap(err, "connect to redis")
	}
	app.Log.Info().Msg("connected to redis")
	return nil
}

func (app *App) Close() error {
	if err := app.Store.Close(); err != nil {
		return errors.Wrap(err, "close postgress connection")
	}

	if err := app.Redis.Close(); err != nil {
		return errors.Wrap(err, "close redis connection")
	}

	return nil
}
