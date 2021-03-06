package usecase

import (
	"fmt"
	"url-shortener/internal/domain"
	"url-shortener/internal/exception"
	"url-shortener/internal/service/utils"
)

type ShortenUrlUsecase struct {
	generator domain.UrlIdGenerator
	storage   domain.Storage
	cache     domain.Cache
	logger    domain.Logger
}

func NewShortenUrlUsecase(generator domain.UrlIdGenerator, storage domain.Storage, cache domain.Cache, logger domain.Logger) *ShortenUrlUsecase {
	return &ShortenUrlUsecase{generator: generator, storage: storage, cache: cache, logger: logger}
}

func (s *ShortenUrlUsecase) Process(shortDomainUrl, longUrl, expireAt string) (string, string, error) {
	urlId := s.generator.NewId()

	expireTime, err := utils.ToExpireSeconds(expireAt)
	if err != nil {
		s.logger.Debugf("expireTime format is invalid. expireTime: %s", expireAt)
		return "", "", exception.ExpireTimeIsNotValid
	}

	err = s.storage.Store(longUrl, urlId, expireAt)
	if err != nil {
		s.logger.Debugf("store url error. longUrl: %s, urlId: %s, expireAt: %s", longUrl, urlId, expireAt)
		s.logger.Errorf("store url error. message: %s", err.Error())
		return "", "", exception.ServerError
	}

	err = s.cache.Set(urlId, longUrl, expireTime)
	if err != nil {
		s.logger.Debugf("cache url error. longUrl: %s, urlId: %s, expireAt: %s", longUrl, urlId, expireAt)
		s.logger.Errorf("cache url error. message: %s", err.Error())
		return "", "", exception.ServerError
	}

	shortUrl := fmt.Sprintf("http://%s/%s", shortDomainUrl, urlId)
	return urlId, shortUrl, nil
}
