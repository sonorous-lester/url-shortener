package validator

import (
	"regexp"
)

type UrlValidator struct {
	reg *regexp.Regexp
}

func NewUrlValidator() UrlValidator {
	reg, err := regexp.Compile("^(?:http(s)?:\\/\\/)?[\\w.-]+(?:\\.[\\w\\.-]+)+[\\w\\-\\._~:/?#[\\]@!\\$&'\\(\\)\\*\\+,;=.]+$")
	if err != nil {
		panic(err)
	}
	return UrlValidator{reg: reg}
}

func (v UrlValidator) Valid(unchecked string) bool {
	return v.reg.MatchString(unchecked)
}
