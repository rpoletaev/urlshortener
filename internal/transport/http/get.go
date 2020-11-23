package http

import (
	"net/http"
	"urlshortener/internal/service"

	"github.com/gorilla/mux"
)

func (api *Api) getLink(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash, ok := vars["hash"]
	if !ok || len(hash) == 0 {
		http.Error(wr, "unknown url", http.StatusBadRequest)
		return
	}

	req := service.GetSourceLinkRequest{Hash: hash}
	resp, err := api.Svc.GetSourceLink(r.Context(), req)
	if err != nil {
		http.Error(wr, err.Error(), errorCode(err))
		return
	}

	http.Redirect(wr, r, resp.Link, http.StatusMovedPermanently)
}
