package internal

type Link struct {
	ID     uint
	Source string
}

type Store interface {
	Create(link string) (uint, error)
	Get(id uint) (Link, error)
}

type HashCodec interface {
	Encode(id uint) string
	Decode(hash string) uint
}

type Cache interface {
	Set(key, val string) error
	Get(key string) (string, error)
}
