package aufgabe2

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	Id   string
	Next *LinkedList
}

func MakeLinkedList(Ids ...string) *LinkedList {
	head := &LinkedList{"", nil}
	for i := len(Ids) - 1; i >= 0; i-- {
		newElement := &LinkedList{Ids[i], head}
		head = newElement
	}
	return head
}

func (list *LinkedList) Empty() bool {
	return list.Next == nil
}

func (list *LinkedList) String() string {
	resultStrings := make([]string, 0)
	for current := list; !current.Empty(); current = current.Next {
		resultStrings = append(resultStrings, current.Id)
	}
	return fmt.Sprintf("[%s]", strings.Join(resultStrings, " "))
}
