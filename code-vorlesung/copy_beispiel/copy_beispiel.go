package copybeispiel

import "fmt"

func CopyDemoNaiv() {
	l1 := []int{1, 2, 3, 4, 5}

	// "Naives" Kopieren einer Liste
	l2 := l1
	fmt.Println(l1)
	fmt.Println(l2)

	// Ändern von l2
	l2[0] = 55555
	fmt.Println(l2)
	fmt.Println(l1)
}

func CopyDemoCopy() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{0, 0, 0, 0, 0}

	fmt.Println(l1)
	fmt.Println(l2)

	// Verwende copy zum Kopieren statt :=
	// Ändere danach ein Element der Liste.
	copy(l2, l1)
	l2[1] = 55555
	fmt.Println(l2)
	fmt.Println(l1)

}
