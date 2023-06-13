package trieslabora

import "testing"

func TestElement_HasValue(t *testing.T) {
	t1 := New()

	if t1.HasValue() {
		t.Error("New element is not empty!")
	}
}

func TestElement_SetValue(t *testing.T) {
	t1 := New()

	t1.SetValue("Hello World")

	if !t1.HasValue() {
		t.Error("Element is empty!")
	}
}

func TestElement_AddValueForPath(t *testing.T) {
	t1 := New()

	t1.AddValueForPath("aaaab", "consetetur")

	node := t1.Left.Left.Left.Left.Right

	if !node.HasValue() {
		t.Error("Element is empty!")
	}
	if node.Value != "consetetur" {
		t.Errorf("element has value %s, but expected %s.", node.Value, "consetetur")
	}
}
