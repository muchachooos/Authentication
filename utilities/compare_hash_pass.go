package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(hashedPass, pass string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	if err != nil {
		return err
	}

	return nil
}
