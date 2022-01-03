package entities

import "testing"

func TestValidateEntityId(t *testing.T) {
	// Invalid ID, not formed correctly
	err := ValidateEntityId("abc123")
	checkError(t, ErrInvalidEntityId, err)

	// Invalid ID, uuid parse error
	err = ValidateEntityId("note.abc123")
	checkError(t, ErrInvalidEntityId, err)

	// Valid ID
	bag := NewCoinBag()
	err = ValidateEntityId(bag.EntityId())
	checkError(t, nil, err)
}
