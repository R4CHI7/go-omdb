package omdb

import (
	"testing"
)

func TestSearchByTitle(t *testing.T) {
	searchTerm := "Into The Wild"
	_, err := SearchByTitle(searchTerm)
	if err != nil {
		t.Errorf("Expected error to be nil. Got: %s", err.Error())
	}
}
