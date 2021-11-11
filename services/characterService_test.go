package services

import (
	"dnd-assistant/entities"
	"errors"
	"testing"
)

type MockedRepo struct {
}

func (m MockedRepo) CreateCharacter(_ *entities.Character) error {
	return nil
}

func (m MockedRepo) GetCharacter(id string) (*entities.Character, error) {
	c := entities.NewCharacter("Barry Allen", "flash")
	c.MaxHp = 10
	c.Ticker = 4
	if id == "none" {
		return nil, errors.New("no character found")
	}

	if id == "ticker 0" {
		c.Ticker = 0
		return c, nil
	}

	if id == "flash" {
		return c, nil
	}

	return c, nil
}

func (m MockedRepo) UpdateCharacter(_ *entities.Character) error {
	return nil
}

func (m MockedRepo) DeleteCharacter(_ string) error {
	return nil
}

func (m MockedRepo) RestoreCharacter(_ string, _ uint) (*entities.Character, error) {
	return entities.NewCharacter("Barry Allen", "flash"), nil
}

func (m MockedRepo) CharacterIds() ([]string, error) {
	return []string{}, nil
}

var repo = MockedRepo{}

func TestCharacterService_Create(t *testing.T) {

	service := NewCharacterService(repo)
	c := entities.NewCharacter("Bruce Wayne", "")
	err := service.Create(c)
	if err == nil {
		t.Error("failed to catch bad data")
	}
	c.Id = "bruce"
	c.MaxHp = 10
	err = service.Create(c)
	if err != nil {
		t.Error(err)
	}
}

func TestCharacterService_Get(t *testing.T) {
	service := NewCharacterService(repo)
	_, _ = service.Get("bruce")
}

func TestCharacterService_Update(t *testing.T) {
	service := NewCharacterService(repo)
	c := entities.NewCharacter("Clark Kent", "superman")
	_ = service.Update(c)
}

func TestCharacterService_Delete(t *testing.T) {
	service := NewCharacterService(repo)
	_ = service.Delete("flash")
}

func TestCharacterService_Undo(t *testing.T) {
	service := NewCharacterService(repo)
	_, err := service.Undo("none")
	if err == nil {
		t.Error("failed to catch none character")
	}
	_, err = service.Undo("ticker 0")
	if err == nil {
		t.Error("failed to catch ticker 0")
	}
	_, _ = service.Undo("flash")

}

func TestCharacterService_GetIds(t *testing.T) {
	service := NewCharacterService(repo)
	_, _ = service.GetIds()
}
