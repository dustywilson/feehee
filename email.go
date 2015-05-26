package feehee

import (
	"errors"
	"net"
	"regexp"
)

type EmailAddress []byte

func NewEmailAddress(email string) (EmailAddress, error) {
	err := testEmailAddress(email)
	if err != nil {
		return nil, err
	}
	return EmailAddress([]byte(email)), nil
}

var emailPattern = regexp.MustCompile(`^[^@]+@([^@]+\.[^@]+)$`)

func testEmailAddress(email string) error {
	if domains := emailPattern.FindStringSubmatch(email); len(domains) > 1 {
		domain := domains[1]
		_, err := net.LookupMX(domain)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Malformed email address")
}

func (e *EmailAddress) Update(email string) error {
	err := testEmailAddress(email)
	if err != nil {
		return err
	}
	*e = []byte(email)
	return nil
}

func (e EmailAddress) String() string {
	return string(e)
}
