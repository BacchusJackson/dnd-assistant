package services

import "dnd-assistant/entities"

type NoteRepo interface {
	CreateNote(n *entities.Note) error
	GetNotes() ([]entities.Note, error)
	UpdateNote(n *entities.Note) (entities.Note, error)
	DeleteNote(id string) error
}

type NoteService struct {
	repo NoteRepo
}

func NewNoteService(repo NoteRepo) *NoteService {
	return &NoteService{repo}
}

// Create a new note in database
func (s NoteService) Create(n *entities.Note) error {
	err := n.Valid()
	if err != nil {
		return err
	}
	return s.repo.CreateNote(n)
}

// GetAll notes from repo
func (s NoteService) GetAll() ([]entities.Note, error) {
	return s.repo.GetNotes()
}

// Update a specific note in repo
func (s NoteService) Update(n *entities.Note) (entities.Note, error) {
	return s.repo.UpdateNote(n)
}

// Delete a specific note from repo
func (s NoteService) Delete(id string) error {
	return s.repo.DeleteNote(id)
}
