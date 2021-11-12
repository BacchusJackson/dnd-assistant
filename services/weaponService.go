package services

import "dnd-assistant/entities"

type WeaponRepo interface {
	CreateWeapon(w *entities.Weapon) error
	GetWeapons() ([]entities.Weapon, error)
	UpdateWeapon(w *entities.Weapon) (*entities.Weapon, error)
	DeleteWeapon(id string) error
}

type WeaponService struct {
	repo WeaponRepo
}

func NewWeaponService(repo WeaponRepo) *WeaponService {
	return &WeaponService{repo: repo}
}

func (s WeaponService) Create(w *entities.Weapon) error {
	err := w.Valid()
	if err != nil {
		return err
	}
	return s.repo.CreateWeapon(w)
}

func (s WeaponService) GetAll() ([]entities.Weapon, error) {
	return s.repo.GetWeapons()
}