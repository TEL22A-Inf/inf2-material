package sortieren

func BubbleSort(list []int) {

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func BubbleUp(list []int) bool {
	swapped := false
	//for j := 0; j < len(list)-1; j++ {
	for j := range list[:len(list)-1] {
		if list[j] > list[j+1] {
			list[j], list[j+1] = list[j+1], list[j]
			swapped = true
		}
	}
	return swapped
}

func BubbleSort2(list []int) {
	// for i := 0; i < len(list); i++ {
	for range list {
		if !BubbleUp(list) {
			return
		}
	}
}
func BubbleSort3(list []int) {
	for BubbleUp(list) {
	}
}
