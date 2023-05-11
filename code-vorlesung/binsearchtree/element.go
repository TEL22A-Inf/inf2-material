package binsearchtree

type Element struct {
	Value int
	Left  *Element
	Right *Element
}

// NewElement erzeugt einen neuen leeren Knoten.
func NewElement() *Element {
	return &Element{}
}

// IsEmpty liefert true, falls der Knoten leer ist.
func (e Element) IsEmpty() bool {
	return e.Left == nil && e.Right == nil
}

// SetValue erwartet einen Wert und schreibt ihn in den Knoten.
// Erzeugt neue Kinder, falls der Knoten leer ist.
func (e *Element) SetValue(value int) {
	if e.IsEmpty() {
		e.Left = NewElement()
		e.Right = NewElement()
	}
	e.Value = value
}

// Insert fügt ein neues Element in den Baum ein.
func (e *Element) Insert(value int) {
	if e.IsEmpty() {
		e.SetValue(value)
		return
	}
	if value < e.Value {
		e.Left.Insert(value)
	} else {
		e.Right.Insert(value)
	}
}
