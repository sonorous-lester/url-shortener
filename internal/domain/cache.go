package domain

type Cache interface {
	Set(key, value string, expireTime int) error
	Get(key string) (string, error)
}
