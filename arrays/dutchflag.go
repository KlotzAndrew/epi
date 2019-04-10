package arrays

func DutchFlagPartition(pivot int, colors []int) []int {
	smaller, equal, larger := 0, 0, len(colors)-1
	target := colors[pivot]

	for equal < larger {
		if colors[equal] < target {
			colors[smaller], colors[equal] = colors[equal], colors[smaller]
			smaller++
			equal++
		} else if colors[equal] > target {
			colors[larger], colors[equal] = colors[equal], colors[larger]
			larger--
		} else {
			equal++
		}
	}
	return colors
}
