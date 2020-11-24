package service

import (
	"context"
	"net/url"
	"urlshortener/internal"

	"github.com/pkg/errors"
)

type Config struct {
	Domain string `envconfig:"DOMAIN"`
}

type Service struct {
	Config
	Store internal.Store
	Cache internal.Cache
	Codec internal.HashCodec
}

type CreateLinkRequest struct {
	Link string
}

type CreateLinkResponse struct {
	ShortLink string
}

func (s *Service) CreateShortLink(cxt context.Context, req CreateLinkRequest) (CreateLinkResponse, error) {

	url, err := url.Parse(req.Link)
	if err != nil {
		return CreateLinkResponse{}, errors.Wrap(internal.ErrBadRequest, err.Error())
	}
	if len(url.Scheme) == 0 {
		return CreateLinkResponse{}, errors.Wrap(internal.ErrBadRequest, "empty scheme")
	}
	id, err := s.Store.Create(req.Link)
	if err != nil {
		return CreateLinkResponse{}, err
	}

	hash := s.Codec.Encode(id)
	resp := CreateLinkResponse{
		ShortLink: s.Domain + "/" + hash,
	}
	return resp, nil
}

type GetSourceLinkRequest struct {
	Hash string
}

type GetSourceLinkResponse struct {
	Link string
}

func (s *Service) GetSourceLink(ctx context.Context, req GetSourceLinkRequest) (GetSourceLinkResponse, error) {
	source, err := s.Cache.Get(req.Hash)
	if err == nil {
		return GetSourceLinkResponse{
			Link: source,
		}, nil
	} else {
		// TODO: log error
	}

	id, err := s.Codec.Decode(req.Hash)
	if err != nil {
		return GetSourceLinkResponse{}, err
	}

	link, err := s.Store.Get(id)
	if err != nil {
		return GetSourceLinkResponse{}, err
	}

	go func() {
		if err := s.Cache.Set(req.Hash, link.Source); err != nil {
			// TODO: log error
		}
	}()

	return GetSourceLinkResponse{
		Link: link.Source,
	}, nil
}
