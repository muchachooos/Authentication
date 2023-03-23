package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(hashedPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}
