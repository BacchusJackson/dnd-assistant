package services

import (
	"dnd-assistant/entities"
	"testing"
)

type mockedRepo struct { }

func (m mockedRepo) CreateWeapon(w *entities.Weapon) error {
	return nil
}

func (m mockedRepo) GetWeapons() ([]entities.Weapon, error) {
	weapons := []entities.Weapon{
		*entities.NewWeapon("club", "1d4", entities.Light),
		*entities.NewWeapon("sword", "1d6", entities.Finesse),
	}
	return weapons, nil
}

func (m mockedRepo) UpdateWeapon(w *entities.Weapon) (*entities.Weapon, error) {
	panic("implement me")
}

func (m mockedRepo) DeleteWeapon(id string) error {
	panic("implement me")
}

var weaponService = NewWeaponService(mockedRepo{})

func TestWeaponService_Create(t *testing.T) {
	weapon := entities.NewWeapon("club", "1d4 Damage", entities.Light)
	weapon.Id = ""
	err := weaponService.Create(weapon)
	if err == nil {
		t.Error("failed to catch missing id")
	}
	weapon.Id = "abc123"
	err = weaponService.Create(weapon)
	if err == nil {
		t.Error("failed to catch malformed id")
	}
	weapon = entities.NewWeapon("club", "1d4 Damage", entities.Light)
	err = weaponService.Create(weapon)
	if err != nil {
		t.Error(err)
	}
}

func TestWeaponService_GetAll(t *testing.T) {
	_, err := weaponService.GetAll()
	if err != nil {
		t.Error(err)
	}
}