//+build wireinject

package main

import (
	"urlshortener/internal"
	"urlshortener/internal/hashids"
	"urlshortener/internal/inmem"
	"urlshortener/internal/postgres"
	"urlshortener/internal/service"
	"urlshortener/internal/transport/http"

	"github.com/google/wire"
	"github.com/rs/zerolog"
)

func providePostgres(c *Config) *postgres.Store {
	return &postgres.Store{
		Config: c.PG,
	}
}

func provideAPIConfig(c *Config) *http.Config {
	return c.HTTP
}

func provideServiceConfig(c *Config) service.Config {
	return *c.Service
}

func provideHashDecoder(c *Config) hashids.Decoder {
	return hashids.New(*c.HashCodec)
}

func provideCache(c *Config) *inmem.Cache {
	return inmem.New(*c.Inmem)
}

func provideApp(logger zerolog.Logger, c *Config) *App {
	wire.Build(
		providePostgres,
		provideAPIConfig,
		provideServiceConfig,
		provideHashDecoder,
		provideCache,
		wire.Struct(new(App), "*"),
		wire.Struct(new(http.Api), "*"),
		wire.Struct(new(service.Service), "*"),
		wire.Bind(new(internal.Store), new(*postgres.Store)),
		wire.Bind(new(internal.Cache), new(*inmem.Cache)),
		wire.Bind(new(internal.HashCodec), new(hashids.Decoder)),
	)

	return &App{}
}
