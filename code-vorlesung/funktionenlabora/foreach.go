package funktionenlabora

func ForEach[T any](list []T, f func(T)) {
	for _, v := range list {
		f(v)
	}
}

func Sum(list []int) int {
	sum := 0
	f := func(v int) {
		sum += v
	}

	ForEach(list, f)

	return sum
}
