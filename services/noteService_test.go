package services

import (
	"dnd-assistant/entities"
	"errors"
	"testing"
)

type fakeRepo struct{}

func (f fakeRepo) CreateNote(n *entities.Note) error {
	if n.Id == "fail" {
		return errors.New("expected")
	}
	return nil
}

func (f fakeRepo) GetNotes() ([]entities.Note, error) {
	return []entities.Note{
		entities.NewNote("Note 1"),
		entities.NewNote("Note 2"),
	}, nil
}

func (f fakeRepo) UpdateNote(_ *entities.Note) (entities.Note, error) {
	return fakeNote, nil
}

func (f fakeRepo) DeleteNote(_ string) error {
	return nil
}

var fakeNote = entities.NewNote("My test note")
var noteService = NewNoteService(fakeRepo{})

func TestNoteService_Create(t *testing.T) {
	note := entities.NewNote("My second test note")
	note.Id = ""
	err := noteService.Create(&note)
	if err == nil {
		t.Error("failed to catch no id")
	}
	_ = noteService.Create(&fakeNote)
}

func TestNoteService_GetAll(t *testing.T) {
	_, _ = noteService.GetAll()
}

func TestNoteService_Update(t *testing.T) {
	_, _ = noteService.Update(&fakeNote)
}

func TestNoteService_Delete(t *testing.T) {
	_ = noteService.Delete("flash")
}
