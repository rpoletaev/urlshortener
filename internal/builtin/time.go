package builtin

import (
	"time"
)

type TimeFunc struct{}

func (t TimeFunc) Now() time.Time {
	return time.Now()
}
