package funktionenlabora

import "sort"

func SortAddressAscendingByStreet(list []Address) {

	less := func(i, j int) bool {
		return list[i].Street < list[j].Street
	}

	sort.Slice(list, less)
}

func SortAddressDescendingByNumber(list []Address) {

	less := func(i, j int) bool {
		return list[i].Number > list[j].Number
	}

	sort.Slice(list, less)
}
