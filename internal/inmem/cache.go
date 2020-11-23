package inmem

import (
	"sync"
	"urlshortener/internal"
)

type Config struct {
	InitMapLength int `envconfig:"CACHE_LEN"`
}

func New(c Config) *Cache {
	return &Cache{
		kv: make(map[string]string, c.InitMapLength),
	}
}

type Cache struct {
	kv map[string]string
	mu sync.RWMutex
}

func (c *Cache) get(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.kv[key]
	if !ok {
		return "", internal.ErrNotFound
	}
	return val, nil
}

func (c *Cache) set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.kv[key] = value
	return nil
}

func (c *Cache) Set(key, val string) error {
	return c.set(key, val)
}

func (c *Cache) Get(key string) (string, error) {
	return c.get(key)
}
