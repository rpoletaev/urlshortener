package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"urlshortener/internal/service"
)

type addLinkRequest struct {
	Link string
}

type addLinkResponse struct {
	ShortLink string
}

func (api *Api) addLink(wr http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(wr, "unable to read request", http.StatusBadRequest)
		return
	}

	r.Body.Close()

	req := addLinkRequest{}
	if err := json.Unmarshal(raw, &req); err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := api.Svc.CreateShortLink(r.Context(), service.CreateLinkRequest{Link: req.Link})
	if err != nil {
		http.Error(wr, err.Error(), errorCode(err))
		return
	}

	if err := writeJson(wr, http.StatusCreated, addLinkResponse{ShortLink: resp.ShortLink}); err != nil {
		// TODO: log error
	}
}
