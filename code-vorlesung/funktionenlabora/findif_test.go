package funktionenlabora

import "fmt"

func ExampleFindIf() {

	l1 := []int{15, 2, 42, 27, 35, 56, 103, 38}

	r1 := FindIfGreater50(l1)
	r2 := FindIfOdd(l1)
	r3 := FindIfBetween20And50(l1)
	r4 := FindIfBetween(l1, 45, 60)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

	// Output:
	// 5
	// 0
	// 2
	// 5
}
