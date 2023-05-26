package funktionenlabora

import "fmt"

func ExampleFindIf_aufgaben() {
	l2 := []int{7, 14, 21, 28, 30, 35, 42, 49, 56, 63, 70, 77, 84, 91, 98, 105}
	l3 := []string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"}
	l4 := []Address{
		{"Hauptstrasse", 1, "1234"},
		{"Nebenstrasse", 42, "5678"},
		{"Dritte Strasse", 123, "1234"},
	}

	r5 := FindIf60IsMultiple(l2)
	r6 := FindIfStartsWithDEF(l3)
	r7 := FindIfContains(l3, "qr")
	r8 := FindIfZipIs1234(l4)
	r9 := FindIfNumberIs42(l4)

	fmt.Println(r5)
	fmt.Println(r6)
	fmt.Println(r7)
	fmt.Println(r8)
	fmt.Println(r9)

	// Output:
	// 4
	// 1
	// 5
	// 0
	// 1
}
