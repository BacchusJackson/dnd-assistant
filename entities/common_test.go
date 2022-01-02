package entities

import "testing"

func checkError(t *testing.T, expected interface{}, got interface{}) {
	if got != expected {
		t.Errorf("expected %v ... got %v\n", expected, got)
	}
}
