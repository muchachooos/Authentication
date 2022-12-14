package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(pass string) (string, error) {

	cost := 10

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}
