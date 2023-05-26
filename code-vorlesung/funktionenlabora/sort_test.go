package funktionenlabora

import "fmt"

func ExampleSortAddressAscendingByStreet() {
	l1 := []Address{
		{"Hauptstrasse", 42, "1234"},
		{"Nebenstrasse", 1, "5678"},
		{"Dritte Strasse", 123, "1234"},
	}
	SortAddressAscendingByStreet(l1)
	for _, v := range l1 {
		fmt.Println(v)
	}

	// Output:
	// {Dritte Strasse 123 1234}
	// {Hauptstrasse 42 1234}
	// {Nebenstrasse 1 5678}
}

func ExampleSortAddressDescendingByNumber() {
	l1 := []Address{
		{"Hauptstrasse", 42, "1234"},
		{"Nebenstrasse", 1, "5678"},
		{"Dritte Strasse", 123, "1234"},
	}
	SortAddressDescendingByNumber(l1)
	for _, v := range l1 {
		fmt.Println(v)
	}

	// Output:
	// {Dritte Strasse 123 1234}
	// {Hauptstrasse 42 1234}
	// {Nebenstrasse 1 5678}
}
