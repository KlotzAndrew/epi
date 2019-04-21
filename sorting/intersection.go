package sorting

func Intersect(a, b []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if (a[i] == b[j]) && (i == 0 || a[i] != a[i-1]) {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		}
	}
	return result
}
