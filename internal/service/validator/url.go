package validator

import (
	"regexp"
)

type UrlValidator struct {
	reg *regexp.Regexp
}

func NewUrlValidator() (UrlValidator, error) {
	reg, err := regexp.Compile("^(?:http(s)?:\\/\\/)?[\\w.-]+(?:\\.[\\w\\.-]+)+[\\w\\-\\._~:/?#[\\]@!\\$&'\\(\\)\\*\\+,;=.]+$")
	if err != nil {
		return UrlValidator{}, err
	}
	return UrlValidator{reg: reg}, nil
}

func (v UrlValidator) Valid(unchecked string) bool {
	return v.reg.MatchString(unchecked)
}
