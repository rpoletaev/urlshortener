// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/rs/zerolog"
	"urlshortener/internal/builtin"
	"urlshortener/internal/hashids"
	"urlshortener/internal/inmem"
	"urlshortener/internal/postgres"
	"urlshortener/internal/redis"
	"urlshortener/internal/service"
	"urlshortener/internal/transport/http"
)

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func provideApp(logger zerolog.Logger, c *Config) *App {
	store := providePostgres(c)
	config := provideAPIConfig(c)
	serviceConfig := provideServiceConfig(c)
	cache := provideCache(c)
	decoder := provideHashDecoder(c)
	backend := provideRedis(c)
	timeFunc := pnovideTimeFunc()
	serviceService := &service.Service{
		Config:   serviceConfig,
		Store:    store,
		Cache:    cache,
		Codec:    decoder,
		Hll:      backend,
		TimeFunc: timeFunc,
	}
	realIPExtractor := provideIPExtractor()
	api := &http.Api{
		Config:      config,
		Svc:         serviceService,
		IpExtractor: realIPExtractor,
	}
	app := &App{
		Store: store,
		API:   api,
		Redis: backend,
		Log:   logger,
	}
	return app
}

// wire.go:

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
