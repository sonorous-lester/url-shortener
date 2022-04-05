package urld

import "github.com/lithammer/shortuuid"

type Generator struct {
}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewId() string {
	return shortuuid.New()
}
