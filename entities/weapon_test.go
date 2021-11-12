package entities

import "testing"

func TestWeapon_Serialization(t *testing.T) {
	weapon := NewWeapon("Sword", "1d6 Damage", Finesse, Light)

	jsonBytes, err := weapon.Marshal()

	if err != nil {
		t.Error(err)
	}
	var weapon2 Weapon

	err = weapon2.Unmarshal(jsonBytes)
	if err != nil {
		t.Error(err)
	}

	t.Log(weapon2)

	var weapon3 Weapon

	err = weapon3.Unmarshal([]byte{})

	if err == nil {
		t.Error("failed to catch json error")
	}

	weapon4 := Weapon{
		Name:        "Sword",
		Description: "1d6 damage",
	}

	weapon4.PropertiesString()
}
