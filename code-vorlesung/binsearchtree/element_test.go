package binsearchtree

import "testing"

func TestElement_NewElement(t *testing.T) {
	root := NewElement()

	if !root.IsEmpty() {
		t.Error("Neues Element sollte leer sein")
	}
}
