package hashids

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDecodeEncode(t *testing.T) {
	tests := []uint{1, 2, 0, 36, 319, 5319}

	c := Config{
		Salt:   "mysuppersalt",
		MinLen: 3,
	}

	codec := New(c)

	for _, v := range tests {
		hash := codec.Encode(v)
		id := codec.Decode(hash)

		assert.Equal(t, v, id)
	}
}
