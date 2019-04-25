package greedy

import "sort"

// HasThreeSum .
func HasThreeSum(arr []int, sum int) bool {
	sort.Ints(arr)
	for _, v := range arr {
		if hasTwoSum(arr, sum-v) {
			return true
		}
	}
	return false
}

func hasTwoSum(arr []int, sum int) bool {
	i, j := 0, len(arr)-1
	for i <= j {
		switch {
		case arr[i]+arr[j] == sum:
			return true
		case arr[i]+arr[j] < sum:
			i++
		case arr[i]+arr[j] > sum:
			j--
		}
	}
	return false
}
