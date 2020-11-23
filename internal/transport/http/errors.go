package http

import (
	"net/http"
	"urlshortener/internal"

	"gopkg.in/errgo.v2/errors"
)

func errorCode(err error) int {
	if errors.Cause(err) == internal.ErrAlreadyExists {
		return http.StatusBadRequest
	}

	if errors.Cause(err) == internal.ErrNotFound {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
