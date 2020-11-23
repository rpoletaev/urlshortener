package main

import (
	"urlshortener/internal/hashids"
	"urlshortener/internal/inmem"
	"urlshortener/internal/postgres"
	"urlshortener/internal/service"
	"urlshortener/internal/transport/http"

	_ "github.com/joho/godotenv/autoload" // preload .env
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	PG        *postgres.Config `envconfig:"DB"`
	HashCodec *hashids.Config  `envconfig:"HASH"`
	Inmem     *inmem.Config    `envconfig:"INMEM"`
	HTTP      *http.Config     `envconfig:"HTTP"`
	Service   *service.Config  `envconfig:"SERVICE"`
}

func LoadConfig(prefix string) (*Config, error) {
	res := &Config{}
	return res, errors.Wrap(envconfig.Process(prefix, res), "parse config")
}
