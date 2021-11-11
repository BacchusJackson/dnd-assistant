package entities

import "testing"

func TestNote_Serialization(t *testing.T) {
	note := NewNote("Some note")
	t.Log(note)

	jsonBytes, err := note.Marshal()

	if err != nil {
		t.Error(err)
	}

	var note2 Note

	err = note2.Unmarshal(jsonBytes)

	if err != nil {
		t.Error(err)
	}

	var note3 Note

	err = note3.Unmarshal([]byte{})

	if err == nil {
		t.Error("failed to catch json error")
	}

}

func TestNote_Valid(t *testing.T) {
	note := Note{}
	_, err := note.Marshal()
	t.Log(err)
	if err == nil {
		t.Error("failed to catch errors")
	}
}
