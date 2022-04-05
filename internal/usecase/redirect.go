package usecase

import (
	"fmt"
	"url-shortener/internal/domain"
	"url-shortener/internal/exception"
)

type RedirectUsecase struct {
	cache domain.Cache
}

func NewRedirectUsecase(cache domain.Cache) *RedirectUsecase {
	return &RedirectUsecase{cache: cache}
}

func (s *RedirectUsecase) Process(key string) (string, error) {
	url, err := s.cache.Get(key)
	if err != nil {
		fmt.Println(err)
		return "", exception.ServerError
	}
	return url, nil
}
