package aufgabe2

import "fmt"

func ExampleLinkedList_Reverse_nonEmptyList() {
	l1 := MakeLinkedList("A", "B", "C")

	fmt.Println(l1.Reverse())

	// Output:
	// [C B A]
}

func ExampleLinkedList_Reverse_singleElement() {
	l1 := MakeLinkedList("A")

	fmt.Println(l1.Reverse())

	// Output:
	// [A]
}

func ExampleLinkedList_Reverse_emptyList() {
	l1 := MakeLinkedList()

	fmt.Println(l1.Reverse())

	// Output:
	// []
}
