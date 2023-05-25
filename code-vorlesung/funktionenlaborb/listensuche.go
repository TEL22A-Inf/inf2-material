package funktionenlaborb

func isLess5(i int) bool {
	return i < 5
}

func FindFirstLess5(list []int) int {
	return FindFirst(list, isLess5)
	/*
		for i, v := range list {
			if isLess5(v) {
				return i
			}
		}
		return -1
	*/
}

func FindFirstEven(list []int) int {

	isEven := func(v int) bool {
		return v%2 == 0
	}

	return FindFirst(list, isEven)
}

func FindFirstBetween2and15(list []int) int {
	return FindFirstBetween(list, 2, 15)
	/*
		return FindFirst(
			list,
			func(v int) bool { return v >= 2 && v <= 15 },
		)
	*/
}

func FindFirstBetween(list []int, a, b int) int {

	pred := func(v int) bool { return v >= a && v <= b }

	return FindFirst(
		list,
		pred,
	)
}

func FindFirst(list []int, predicate func(int) bool) int {
	for i, v := range list {
		if predicate(v) {
			return i
		}
	}
	return -1
}

/* Schreiben Sie Funktionen, die...

* Wie Less5 funktionieren, aber die Grenze als Parameter erwarten.
* Aus einer Liste von Strings den ersten liefern, der ...
  * weniger als 3 Buchstaben hat
  * mit "abc" anfängt
  * mindestens einen Großbuchstaben enthält.
* Aus einer Liste von Personen die erste liefert, die ...
  * älter als 42 Jahre ist.
  * deren Name mit "C" beginnt.

*/

type Person struct {
	Name string
	Age  int
}

func FilterByAgeLess(list []Person, maxAge int) []Person {
	result := []Person{}
	for _, p := range list {
		if p.Age <= maxAge {
			result = append(result, p)
		}
	}
	return result
}
