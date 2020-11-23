package http

import (
	"encoding/json"
	"net/http"
)

func writeJson(wr http.ResponseWriter, status int, val interface{}) error {
	wr.WriteHeader(status)
	return json.NewEncoder(wr).Encode(val)
}
