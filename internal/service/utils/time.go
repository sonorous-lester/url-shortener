package utils

import (
	"errors"
	"time"
)

func ToExpireSeconds(expireAt string) (int, error) {
	expireTime, err := time.Parse(
		time.RFC3339, expireAt)

	if err != nil {
		return 0, errors.New("time format valid")
	}

	now := time.Now().UTC()
	delta := int(expireTime.Sub(now).Seconds())

	if delta < 0 {
		return 0, errors.New("time format valid")
	}
	return delta, nil
}
