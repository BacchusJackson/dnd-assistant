package entities

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"strings"
)

type Entity interface {
	Map() map[string]string
	EntityId() string
}

var ErrInvalidEntityId = errors.New("invalid entity id")

func ValidateEntityId(id string) error {
	parts := strings.Split(id, ".")
	if len(parts) != 2 {
		return ErrInvalidEntityId
	}
	_, err := uuid.Parse(parts[1])
	if err != nil {
		log.Printf("Invalid ID: %s\n", err)
		return ErrInvalidEntityId
	}
	return nil
}
