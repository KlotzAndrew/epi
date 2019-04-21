package search

func FirstK(arr []int, target int) int {
	left, right, result := 0, len(arr)-1, -1

	for left < right {
		mid := left + ((right - left) / 2)
		switch v := arr[mid]; {
		case v > target:
			right = mid - 1
		case v == target:
			result = mid
			right = mid - 1
		case v < target:
			left = mid + 1
		}
	}
	return result
}
