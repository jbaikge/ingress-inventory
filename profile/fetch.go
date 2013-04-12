package profile

import (
	"errors"
)

var NotFound = errors.New("Profile not found")

func Fetch(id string) (p *Profile, err error) {
	return
}
