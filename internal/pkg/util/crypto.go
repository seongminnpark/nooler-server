package util

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
)

func SaltAndHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashAndPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
