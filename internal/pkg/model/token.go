package model

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	// TokenString string `json:"tokenString"`
	Exp  int64  `json:"exp"`
	UUID string `json:"uuid"`
}

func (token *Token) Encode() (string, error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["uuid"] = token.UUID
	claims["exp"] = token.Exp

	secret := os.Getenv("TOKEN_KEY")
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (token *Token) Decode(tokenString string) error {
	// Validate the token
	jwtToken, parseErr := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_KEY")), nil
	})

	if parseErr != nil {
		return parseErr
	}

	// There was no error while parsing, but token may still be invalid.
	if !jwtToken.Valid {
		return errors.New("Token not valid")
	}
	if err := token.FromTokenObject(jwtToken); err != nil {
		return err
	}
	// There was no error.
	return nil
}

func (token *Token) FromTokenObject(jwtToken *jwt.Token) error {

	// Extract claims.
	claims := jwtToken.Claims.(jwt.MapClaims)

	// Extract exp.
	// if exp, ok := claims["exp"].(float64); ok {
	// 	token.Exp = exp
	// } else {
	// 	return errors.New("Expiration not valid")
	// }

	// Extract uuid.
	if uuid, ok := claims["uuid"].(string); ok {
		token.UUID = uuid
	} else {
		return errors.New("UUID not valid")
	}

	return nil

}

// func (token *Token) ToTokenObject() jwt.Token {
// 	return
// }
