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

func TestWeapon_Map(t *testing.T) {
	weapon := NewWeapon("Katana", "1d6 light weapon", Finesse, Light)
	weaponMap := weapon.Map()
	t.Log(weaponMap)
}

func TestParseWeapon(t *testing.T) {
	weapon := NewWeapon("Katana", "1d6 light weapon", Finesse, Light)
	weaponMap := weapon.Map()

	weapon2, err := ParseWeapon(weaponMap)
	checkError(t, nil, err)
	checkError(t, weapon.Id, weapon2.Id)
	t.Log(weapon2)
}