package redis

import (
	"time"
	"urlshortener/internal"

	"github.com/gomodule/redigo/redis"
)

const (
	ipPref  = "ip"
	urlPref = "url"
	layout  = "20060102"
)

type StatisticsRepository Backend

func (r *Backend) StatisticsRepository() internal.StatisticsRepository {
	return (*StatisticsRepository)(r)
}

func (r *StatisticsRepository) AddIP(ip string, date time.Time) error {
	return r.addStat(ipPref, date, ip)
}

func (r *StatisticsRepository) AddURL(url string, date time.Time) error {
	return r.addStat(urlPref, date, url)
}

func (r *StatisticsRepository) IPStat(dateFrom, dateTo time.Time) (uint, error) {
	return r.getStat(ipPref, dateFrom, dateTo)
}

func (r *StatisticsRepository) URLStat(dateFrom, dateTo time.Time) (uint, error) {
	return r.getStat(urlPref, dateFrom, dateTo)
}

func (r *StatisticsRepository) addStat(keyPref string, date time.Time, val string) error {
	con := r.Pool.Get()
	defer con.Close()

	key := keyPref + date.Format(layout)
	_, err := con.Do("PFADD", key, val)
	return err
}

func (r *StatisticsRepository) getStat(keyPref string, dateFrom, dateTo time.Time) (uint, error) {
	con := r.Pool.Get()
	defer con.Close()

	keys := []interface{}{}
	for d := dateFrom; d.Before(dateTo); d = d.Add(24 * time.Hour) {
		keys = append(keys, keyPref+d.Format(layout))
	}

	val, err := redis.Uint64(con.Do("PFCOUNT", keys...))
	return uint(val), err
}
