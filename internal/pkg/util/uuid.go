package util

import (
	"github.com/satori/go.uuid"
)

func CreateUUID() (string, error) {
	newUUID, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return newUUID.String(), nil
}
