package exception

import "errors"

var (
	URLIsNotValid        = errors.New("url is not valid")
	ExpireTimeIsNotValid = errors.New("expire time is not valid")
	ServerError          = errors.New("server error")
	IncorrectInput       = errors.New("incorrect input")
)
