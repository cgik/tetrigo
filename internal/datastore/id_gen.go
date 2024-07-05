package datastore

import (
	"errors"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	availableValues = "123456789abcdefghjkmnpqrstuvwxyz"
	length          = 8
)

func GenerateId() (string, error) {
	return gonanoid.Generate(availableValues, length)
}

func Validate(fieldName, id string) error {
	if id == "" {
		return errors.New(fieldName + " must not be empty")
	}

	if len(id) != length {
		return errors.New(fieldName + " incorrect number of characters")
	}

	if strings.Trim(id, availableValues) != "" {
		return errors.New(fieldName + " contains invalid characters")
	}

	return nil
}
