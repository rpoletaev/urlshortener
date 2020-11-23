package hashids

import (
	"urlshortener/internal"

	"github.com/speps/go-hashids"
)

type Config struct {
	Salt   string `envconfig:"SALT"`
	MinLen int    `envconfig:"MIN_LEN"`
}

type Decoder struct {
	hasher *hashids.HashID
}

func New(c Config) Decoder {
	hd := hashids.NewData()
	hd.Salt = c.Salt
	hd.MinLength = c.MinLen
	hasher, _ := hashids.NewWithData(hd)
	return Decoder{hasher: hasher}
}

func (d Decoder) Encode(id uint) string {
	hash, _ := d.hasher.Encode([]int{int(id)})
	return hash
}

func (d Decoder) Decode(hash string) (id uint, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = internal.ErrNotFound
		}
	}()
	ids := d.hasher.Decode(hash)
	return uint(ids[0]), nil
}
