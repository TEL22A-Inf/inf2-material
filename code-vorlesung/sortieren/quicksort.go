package sortieren

func QuickSort(list []int) {
	if len(list) < 2 {
		return
	}

	// Pivot wählen
	pivot := list[0]

	// Liste zerlegen
	left := []int{}
	right := []int{}
	for _, el := range list[1:] {
		if el <= pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	// left und right auf magische Weise sortieren.
	QuickSort(left)
	QuickSort(right)

	// left + pivot + right wieder zusammensetzen.
	result := []int{}
	result = append(result, left...)
	result = append(result, pivot)
	result = append(result, right...)

	copy(list, result)
}
