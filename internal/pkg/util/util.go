package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

func CreateUUID() (string, error) {
	newUUID, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return newUUID.String(), nil
}

func GenerateToken(secret string, claimsMap map[string]interface{}) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// Dump claims from parameters to token.
	for k, v := range claimsMap {
		claims[k] = v
	}

	/* Sign the token with our secret */
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
