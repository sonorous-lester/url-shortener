package domain

type Storage interface {
	Store(longUrl, urlId, expireAt string) error
}
