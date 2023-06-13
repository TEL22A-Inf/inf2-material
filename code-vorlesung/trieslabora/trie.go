package trieslabora

type Element struct {
	Key   string
	Left  *Element
	Right *Element
	Value string
}

func New() *Element {
	return &Element{}
}

func (e Element) HasValue() bool {
	return e.Value != ""
}

func (e *Element) SetValue(newValue string) {
	e.Value = newValue
}

func (e *Element) AddValueForPath(path, value string) {
	if path == "" {
		e.SetValue(value)
		return
	}
	head, tail := path[:1], path[1:]

	if head == "a" {
		if e.Left == nil {
			e.Left = New()
		}
		e.Left.AddValueForPath(tail, value)
	} else {
		if e.Right == nil {
			e.Right = New()
		}
		e.Right.AddValueForPath(tail, value)
	}
}
