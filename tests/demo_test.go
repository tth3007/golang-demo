package tests

import (
	"demo/structs"
	"testing"
)

func TestBookType(t *testing.T) {
	longBook := structs.Book{
		ID:    1,
		Name:  "Demo",
		Pages: 10,
	}

	var bookType string = longBook.GetType()

	if bookType != "NOVEL" {
		t.Errorf("Failed! Expect: %q, Actual: %q", "NOVEL", longBook.GetType())
	}
}
