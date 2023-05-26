package funktionenlabora

func Greater50(v int) bool {
	return v > 50
}

// FindIfGreater50 erwartet eine Liste und liefert die Position des ersten Elements,
// das größer als 50 ist.
func FindIfGreater50(list []int) int {
	return FindIf_int(list, Greater50)
	/*
		for i, v := range list {
			if v > 50 {
				return i
			}
		}
		return -1
	*/
}

// FindIfOdd erwartet eine Liste und liefert die Position des ersten Elements,
// das ungerade ist.
func FindIfOdd(list []int) int {

	isOdd := func(v int) bool {
		return v%2 != 0
	}
	return FindIf_int(list, isOdd)
	/*
		for i, v := range list {
			if v%2 != 0 {
				return i
			}
		}
		return -1
	*/
}

// FindIfBetween20And50 erwartet eine Liste und liefert die Position des ersten Elements,
// das zwischen 20 und 50 liegt.
func FindIfBetween20And50(list []int) int {

	return FindIf_int(
		list,
		func(v int) bool { return v > 20 && v < 50 },
	)

	/*
		for i, v := range list {
			if v > 20 && v < 50 {
				return i
			}
		}
		return -1
	*/
}

// FindIfBetween erwartet eine Liste und zwei Zahlen a,b.
// Sie liefert die Position des ersten Elements,
// das zwischen a und b liegt.
func FindIfBetween(list []int, a, b int) int {

	return FindIf_int(
		list,
		func(v int) bool { return v > a && v < b },
	)
}

/*
func p(v int) bool {
	// return v%2 == 1
	return v > 20 && v < 50
}
*/

// FindIf_int
func FindIf_int(list []int, p func(int) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}
