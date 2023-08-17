package searching

func InterpolationSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high && target >= arr[low] && target <= arr[high] {

		// If array has just a single element
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}

		index := low + ((target - arr[low]) * (high - low) / (arr[high] - arr[low]))

		if arr[index] == target {
			return index
		} else if arr[index] < target {
			low = index + 1
		} else {
			high = index - 1
		}

	}

	return -1
}
