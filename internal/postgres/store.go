package postgres

import (
	"strings"
	"urlshortener/internal"

	"github.com/jmoiron/sqlx"
)

// Config for postgres connection
type Config struct {
	Driver       string `envconfig:"DRIVER"`
	URI          string `envconfig:"URI"`
	MaxOpenConns int    `envconfig:"MAX_OPEN"`
	MaxIdleConns int    `envconfig:"MAX_IDLE"`
}

func New(c *Config) *Store {
	return &Store{Config: c}
}

type Store struct {
	*Config
	db *sqlx.DB
}

func (s *Store) Connect() error {
	db, err := sqlx.Connect(s.Driver, s.URI)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(s.MaxOpenConns)
	db.SetMaxIdleConns(s.MaxIdleConns)
	s.db = db

	return s.Init()
}

func (s *Store) Close() error {
	return s.db.Close()
}

const (
	initQuery = `CREATE TABLE IF NOT EXISTS links (
		id SERIAL PRIMARY KEY,
		source text NOT NULL,
		CONSTRAINT	unique_source UNIQUE(source))`

	createQuery = "INSERT INTO links (source) VALUES ($1) RETURNING id"
	selectQuery = "SELECT * FROM links WHERE id = $1"
)

func (s *Store) Init() error {
	if _, err := s.db.Exec(initQuery); err != nil {
		return err
	}

	return nil
}

func (s *Store) Create(link string) (uint, error) {
	var newId uint
	if err := s.db.QueryRow(createQuery, link).Scan(&newId); err != nil {
		if strings.Contains(err.Error(), "unique") {
			return 0, internal.ErrAlreadyExists
		}
		return 0, err
	}
	return newId, nil
}

func (s *Store) Get(id uint) (internal.Link, error) {
	lnk := internal.Link{}
	if err := s.db.Get(&lnk, selectQuery, id); err != nil {
		return lnk, internal.ErrNotFound
	}
	return lnk, nil
}
