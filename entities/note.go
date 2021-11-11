package entities

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

const TimeFormat = time.RFC822

type Note struct {
	Id   string `json:"id"`
	Body string `json:"body"`
	Date string `json:"date"`
}

func NewNote(body string) Note {
	n := Note{Body: body}
	n.Touch()
	n.Id = uuid.NewString()
	return n
}

// Marshal converts the note into JSON bytes
func (n Note) Marshal() ([]byte, error) {
	err := n.Valid()
	if err != nil {
		return nil, err
	}
	return json.Marshal(n)
}

// Unmarshal parses json data into a note
func (n *Note) Unmarshal(jsonBytes []byte) error {
	err := json.Unmarshal(jsonBytes, &n)
	if err != nil {
		return err
	}
	return n.Valid()
}

// Valid checks if the note has the appropriate fields.
// Returns nil if valid
func (n Note) Valid() error {
	var b strings.Builder
	if n.Id == "" {
		b.WriteString("no id for note\n")
	}
	if n.Id == "" {
		b.WriteString("no character id for note\n")
	}
	_, err := time.Parse(TimeFormat, n.Date)
	if err != nil {
		b.WriteString("bad time format\n")
	}
	if n.Body == "" {
		b.WriteString("no body\n")
	}
	if b.Len() != 0 {
		return errors.New(b.String())
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
