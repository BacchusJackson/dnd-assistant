package repositories

import (
	"dnd-assistant/entities"
	"errors"
)

var ErrEntityNotFound = errors.New("requested entity not found")
var ErrRepoWriteFail = errors.New("repository did not write: Database Error")
var ErrRepoInvalidEntity = errors.New("repository did not write: Invalid Entity")

type Repository interface {
	Write(e entities.Entity) error
	Read(entityId string) error
}
