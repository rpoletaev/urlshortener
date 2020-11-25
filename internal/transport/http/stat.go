package http

import (
	"net/http"
	"strconv"
	"time"
	"urlshortener/internal"
	"urlshortener/internal/service"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (api *Api) getIPStat(wr http.ResponseWriter, r *http.Request) {

	p, err := parseStatPeriod(r)
	if err != nil {
		http.Error(wr, err.Error(), errorCode(err))
		return
	}
	req := service.StatRequest{
		From: p.from,
		To:   p.to,
	}
	resp, err := api.Svc.GeIPtStat(r.Context(), req)
	if err != nil {
		http.Error(wr, err.Error(), errorCode(err))
		return
	}

	writeJson(wr, http.StatusOK, resp)
}

func (api *Api) getURLStat(wr http.ResponseWriter, r *http.Request) {

	p, err := parseStatPeriod(r)
	if err != nil {
		http.Error(wr, err.Error(), errorCode(err))
		return
	}
	req := service.StatRequest{
		From: p.from,
		To:   p.to,
	}
	resp, err := api.Svc.GetURLStat(r.Context(), req)
	if err != nil {
		http.Error(wr, err.Error(), errorCode(err))
		return
	}

	writeJson(wr, http.StatusOK, resp)
}

type period struct {
	from time.Time
	to   time.Time
}

func parseStatPeriod(r *http.Request) (period, error) {
	vars := mux.Vars(r)
	fromv, ok := vars["from"]
	if !ok || len(fromv) == 0 {
		return period{}, errors.Wrap(internal.ErrBadRequest, "period's start must be set")
	}

	fromint, err := strconv.Atoi(fromv)
	if err != nil {
		return period{}, errors.Wrap(internal.ErrBadRequest, "period's start must be unixtime")
	}

	tov, ok := vars["to"]
	if !ok || len(tov) == 0 {
		return period{}, errors.Wrap(internal.ErrBadRequest, "period's end must be set")
	}

	toint, err := strconv.Atoi(tov)
	if err != nil {
		return period{}, errors.Wrap(internal.ErrBadRequest, "period's end must be unixtime")
	}

	return period{
		from: time.Unix(int64(fromint), 0),
		to:   time.Unix(int64(toint), 0),
	}, nil
}
