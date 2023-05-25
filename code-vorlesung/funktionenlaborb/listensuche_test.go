package funktionenlaborb

import "fmt"

func ExampleFindFirst() {
	l1 := []int{17, 22, 45, 13, 2, 10, 25}

	r1 := FindFirstLess5(l1)
	r2 := FindFirstEven(l1)
	r3 := FindFirstBetween2and15(l1)
	r4 := FindFirstBetween(l1, 30, 50)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

	// Output:
	// 4
	// 1
	// 3
	// 2
}
