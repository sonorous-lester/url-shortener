package domain

type UrlValidator interface {
	Valid(unchecked string) bool
}
