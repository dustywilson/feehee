package feehee

import (
	"errors"
	"regexp"
	"strings"
)

type Username []byte

func NewUsername(username string) (Username, error) {
	err := testUsername(username)
	if err != nil {
		return nil, err
	}
	return Username([]byte(username)), nil
}

var usernamePattern = regexp.MustCompile(`^[a-z0-9\-\.]+$`)

func testUsername(username string) error {
	username = strings.ToLower(username)
	if usernamePattern.MatchString(username) {
		return nil
	}
	return errors.New("Malformed username")
}

func (u *Username) Update(username string) error {
	err := testUsername(username)
	if err != nil {
		return err
	}
	*u = []byte(username)
	return nil
}

func (u Username) String() string {
	return string(u)
}
