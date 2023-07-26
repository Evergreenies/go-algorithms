package searching

func LinearSearch(arr []int, target int) int {
	for index, element := range arr {
		if element == target {
			return index
		}
	}
	return -1
}
