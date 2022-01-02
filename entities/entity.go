package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

func ValidId(id string) error {
	parts := strings.Split(id, ".")
	if len(parts) != 2 {

		return errors.New("invalid ID")
	}
	_, err := uuid.Parse(parts[1])

	if err != nil {
		return errors.New("invalid ID")
	}
	return nil
}
