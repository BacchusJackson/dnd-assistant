package entities

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

const TimeFormat = time.RFC822

type Note struct {
	Id   string `json:"id"`
	Body string `json:"body"`
	Date string `json:"date"`
}

func NewNote(body string) *Note {
	n := &Note{Body: body}
	n.Touch()
	n.Id = fmt.Sprintf("note.%s", uuid.NewString())
	return n
}

func ParseNote(m map[string]string) (*Note, error) {
	n := &Note{
		Id:   m["id"],
		Body: m["body"],
		Date: m["date"],
	}
	return n, n.Valid()
}

var ErrInvalidNote = errors.New("invalid note")

// Valid checks if the note has the appropriate fields.
// Returns nil if valid
func (n Note) Valid() error {
	err := ValidId(n.Id)

	if err != nil {
		log.Println("invalid ID")
		return ErrInvalidNote
	}

	_, err = time.Parse(TimeFormat, n.Date)
	if err != nil {
		log.Println("invalid time format")
		return ErrInvalidNote
	}

	return nil
}

// Touch updates the note to the current time
func (n *Note) Touch() {
	n.Date = time.Now().Format(TimeFormat)
}

// String converts the note to a formatted string
func (n Note) String() string {
	return n.Date + " - " + n.Id + "\n" + n.Body + "\n"
}

// Map converts a note to a map
func (n Note) Map() map[string]string {
	return map[string]string{
		"id":   n.GetId(),
		"body": n.Body,
		"date": n.Date,
	}
}

func (n Note) GetId() string {
	return n.Id
}
