//+build wireinject

package main

import (
	"urlshortener/internal"
	"urlshortener/internal/builtin"
	"urlshortener/internal/hashids"
	"urlshortener/internal/inmem"
	"urlshortener/internal/postgres"
	"urlshortener/internal/redis"
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

func provideRedis(c *Config) *redis.Backend {
	return &redis.Backend{
		Config: c.Redis,
	}
}

func provideIPExtractor() http.RealIPExtractor {
	return http.RealIPExtractor{}
}

func pnovideTimeFunc() builtin.TimeFunc {
	return builtin.TimeFunc{}
}

func provideApp(logger zerolog.Logger, c *Config) *App {
	wire.Build(
		providePostgres,
		provideAPIConfig,
		provideServiceConfig,
		provideHashDecoder,
		provideCache,
		provideRedis,
		provideIPExtractor,
		pnovideTimeFunc,
		wire.Struct(new(App), "*"),
		wire.Struct(new(http.Api), "*"),
		wire.Struct(new(service.Service), "*"),
		wire.Bind(new(internal.Store), new(*postgres.Store)),
		wire.Bind(new(internal.Cache), new(*inmem.Cache)),
		wire.Bind(new(internal.HashCodec), new(hashids.Decoder)),
		wire.Bind(new(internal.Hll), new(*redis.Backend)),
		wire.Bind(new(http.IPExtractor), new(http.RealIPExtractor)),
		wire.Bind(new(internal.TimeFunc), new(builtin.TimeFunc)),
	)

	return &App{}
}
