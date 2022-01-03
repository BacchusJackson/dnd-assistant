package entities

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"strings"
)

type Entity interface {
	Map() map[string]string
	EntityKey() string
	Validate() error
}

var ErrInvalidEntity = errors.New("invalid entity")
var ErrInvalidKey = errors.New("invalid entity id")

func Validate(key string) error {
	_, _, err := Unpack(key)
	return err
}

func Unpack(key string) (id string, eType string, err error) {
	parts := strings.Split(key, ".")
	if len(parts) != 2 {
		return "", "", ErrInvalidKey
	}
	_, err = uuid.Parse(parts[1])
	if err != nil {
		log.Printf("Invalid ID ->%s<-: %s\n", id, err)
		return "", "", ErrInvalidKey
	}

	return parts[0], parts[1], nil
}
