package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func GenerateHashPassword(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}
