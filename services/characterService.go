package services

import (
	"dnd-assistant/entities"
	"errors"
)

// CharacterRepo the methods a repository must implement to work with a character
type CharacterRepo interface {
	CreateCharacter(c *entities.Character) error
	GetCharacter(id string) (*entities.Character, error)
	UpdateCharacter(c *entities.Character) error
	DeleteCharacter(id string) error
	RestoreCharacter(id string, ticker uint) (*entities.Character, error)
	CharacterIds() ([]string, error)
}

// CharacterService the object used to manage character data
type CharacterService struct {
	repo CharacterRepo
}

// NewCharacterService returns a pointer to a character noteService used to interact with the repository
func NewCharacterService(repo CharacterRepo) *CharacterService {
	return &CharacterService{repo: repo}
}

// Create add a character to the repo
func (s CharacterService) Create(c *entities.Character) error {
	err := c.Valid()
	if err != nil {
		return err
	}
	return s.repo.CreateCharacter(c)
}

// Get return a character by id
func (s CharacterService) Get(id string) (*entities.Character, error) {
	return s.repo.GetCharacter(id)
}

// Update pass the updated character object
func (s CharacterService) Update(c *entities.Character) error {
	return s.repo.UpdateCharacter(c)
}

// Delete a character by id
func (s CharacterService) Delete(id string) error {
	return s.repo.DeleteCharacter(id)
}

// Undo any changes made to a character by id
func (s CharacterService) Undo(id string) (*entities.Character, error) {
	c, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	if c.Ticker == 0 {
		return nil, errors.New("cannot undo character... no history")
	}

	return s.repo.RestoreCharacter(c.Id, c.Ticker-1)

}

// GetIds returns a slice of character ids in the repo
func (s CharacterService) GetIds() ([]string, error) {
	return s.repo.CharacterIds()
}
