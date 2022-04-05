package domain

type UrlIdGenerator interface {
	NewId() string
}
