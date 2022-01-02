package entities

import "testing"

func TestNote_Valid(t *testing.T) {
	// blank note
	note := &Note{}
	err := note.Valid()
	checkError(t, ErrInvalidNote, err)

	note.Id = "invalid.id"
	err = note.Valid()
	checkError(t, ErrInvalidNote, err)

	// invalid time format
	note = NewNote("content")
	note.Date = "invalid"
	err = note.Valid()
	checkError(t, ErrInvalidNote, err)

	// good note
	note = NewNote("content")
	err = note.Valid()
	checkError(t, nil, err)
}

func TestNote_String(t *testing.T) {
	note := NewNote("content")
	t.Log(note)
}

func TestNote_Map(t *testing.T) {
	note := NewNote("content")
	noteMap := note.Map()
	t.Log(noteMap)
}

func TestNote_GetId(t *testing.T) {
	note := NewNote("content")
	t.Log(note.GetId())
}

func TestParseNote(t *testing.T) {
	note := NewNote("content")
	noteMap := note.Map()
	note2, err := ParseNote(noteMap)
	checkError(t, nil, err)
	checkError(t, note.Id, note2.Id)
}
