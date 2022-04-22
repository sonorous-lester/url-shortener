package usecase

import (
	"url-shortener/internal/domain"
	"url-shortener/internal/exception"
)

type RedirectUsecase struct {
	cache  domain.Cache
	logger domain.Logger
}

func NewRedirectUsecase(cache domain.Cache, logger domain.Logger) *RedirectUsecase {
	return &RedirectUsecase{cache: cache, logger: logger}
}

func (s *RedirectUsecase) Process(key string) (string, error) {
	url, err := s.cache.Get(key)
	if err != nil {
		s.logger.Errorf("get url from cache error. message: %s", err.Error())
		return "", exception.ServerError
	}
	return url, nil
}
