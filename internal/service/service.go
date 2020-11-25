package service

import (
	"context"
	"net/url"
	"time"
	"urlshortener/internal"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Config struct {
	Domain string `envconfig:"DOMAIN"`
}

type Service struct {
	Config
	Store    internal.Store
	Cache    internal.Cache
	Codec    internal.HashCodec
	Hll      internal.Hll
	TimeFunc internal.TimeFunc
	Log      zerolog.Logger
}

func (s *Service) logger(ctx context.Context) *zerolog.Logger {
	l := s.Log.With().Str("component", "service").Logger()
	return &l
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
		s.logger(ctx).Error().Err(err).Str("hash", req.Hash).Msg("on get link from cache")
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
			s.logger(ctx).Error().Err(err).Str("hash", req.Hash).Msg("on save hash to cache")
		}
	}()

	return GetSourceLinkResponse{
		Link: link.Source,
	}, nil
}

type AddIPStatRequest struct {
	IP string
}

type AddURLStatRequest struct {
	URL string
}

func (s *Service) AddIPStat(ctx context.Context, r AddIPStatRequest) error {
	s.logger(ctx).Info().Str("ip", r.IP).Msg("trace ip")
	return s.Hll.StatisticsRepository().AddIP(r.IP, s.TimeFunc.Now())
}

func (s *Service) AddURLStatu(ctx context.Context, r AddURLStatRequest) error {
	s.logger(ctx).Info().Str("url", r.URL).Msg("trace url")
	return s.Hll.StatisticsRepository().AddURL(r.URL, s.TimeFunc.Now())
}

type StatRequest struct {
	From time.Time
	To   time.Time
}

type StatResponse struct {
	Count uint
}

func (s *Service) GeIPtStat(ctx context.Context, r StatRequest) (StatResponse, error) {
	s.logger(ctx).Info().Interface("stat-req", r).Msg("trace stat request")

	count, err := s.Hll.StatisticsRepository().IPStat(r.From, r.To)
	return StatResponse{
		Count: count,
	}, err
}

func (s *Service) GetURLStat(ctx context.Context, r StatRequest) (StatResponse, error) {
	s.logger(ctx).Info().Interface("stat-req", r).Msg("trace stat request")

	count, err := s.Hll.StatisticsRepository().URLStat(r.From, r.To)
	return StatResponse{
		Count: count,
	}, err
}
