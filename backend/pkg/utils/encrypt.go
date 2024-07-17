package utils

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), bcrypt.DefaultCost)
	return string(bytes), err
}

func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(strings.TrimSpace(hashedPassword)), []byte(strings.TrimSpace(password)))
}
