package internal

import "time"

type Link struct {
	ID     uint
	Source string
}

type TimeFunc interface {
	Now() time.Time
}
type Store interface {
	Create(link string) (uint, error)
	Get(id uint) (Link, error)
}

type HashCodec interface {
	Encode(id uint) string
	Decode(hash string) (uint, error)
}

type Cache interface {
	Set(key, val string) error
	Get(key string) (string, error)
}

type StatisticsRepository interface {
	AddIP(ip string, date time.Time) error
	AddURL(url string, date time.Time) error
	IPStat(dateFrom, dateTo time.Time) (uint, error)
	URLStat(dateFrom, dateTo time.Time) (uint, error)
}

type Hll interface {
	StatisticsRepository() StatisticsRepository
}
