package funktionenlabora

func FindIf_String(list []string, p func(string) bool) int {
	return FindIf(list, p)
	/*
		for i, v := range list {
			if p(v) {
				return i
			}
		}
		return -1
	*/
}

func FindIf_Address(list []Address, p func(Address) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

func FindIf[T any](list []T, p func(T) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

func FindIf60IsMultiple(list []int) int {
	// TODO
	return 0
}

func FindIfStartsWithDEF(list []string) int {
	// TODO
	return 0
}

func FindIfContains(list []string, s string) int {
	// TODO
	return 0
}

type Address struct {
	Street string
	Number int
	Zip    string
}

func FindIfZipIs1234(list []Address) int {
	// TODO
	return 0
}

func FindIfNumberIs42(list []Address) int {
	// TODO
	return 0
}
