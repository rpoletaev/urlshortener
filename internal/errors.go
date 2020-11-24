package internal

import "errors"

var ErrAlreadyExists = errors.New("already exists")
var ErrNotFound = errors.New("not found")
var ErrBadRequest = errors.New("bad request")
