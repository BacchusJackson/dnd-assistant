package entities

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestValidateEntityId(t *testing.T) {
	// Invalid ID, not formed correctly
	err := Validate("abc123")
	checkError(t, ErrInvalidKey, err)

	// Invalid ID, uuid parse error
	err = Validate("note.abc123")
	checkError(t, ErrInvalidKey, err)

	// Blank entry
	err = Validate("")
	checkError(t, ErrInvalidKey, err)
	// Validate ID
	id := fmt.Sprintf("coin_bag.%s", uuid.NewString())
	err = Validate(id)
	checkError(t, nil, err)
}
