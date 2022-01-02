package entities

import "testing"

func TestWeapon_String(t *testing.T) {
	weapon := NewWeapon("Katana", "1d6 light weapon", Finesse, Light)
	t.Log(weapon)
	weapon = NewWeapon("katana", "")
	t.Log(weapon)
}

func TestWeapon_Valid(t *testing.T) {
	weapon := NewWeapon("Katana", "1d6 light weapon", Finesse, Light)
	weapon.Id = "bad.id"
	err := weapon.Valid()
	checkError(t, ErrInvalidWeapon, err)

	weapon = NewWeapon("Katana", "1d6 light weapon", Finesse, Light)
	err = weapon.Valid()
	checkError(t, nil, err)
}
