package domain

type ShortenUrlUsecase interface {
	Process(shortDomainUrl, longUrl, expireAt string) (string, string, error)
}

type RedirectUsecase interface {
	Process(key string) (string, error)
}
