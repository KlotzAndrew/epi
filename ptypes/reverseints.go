package ptypes

func ReverseDigits(x int64) int64 {
	sign := int64(1)
	if x < 0 {
		sign = -1
	}

	remaining := uint64(x * sign)

	result := uint64(0)
	for remaining != 0 {
		result = result*10 + remaining%10
		remaining /= 10
	}
	return int64(result) * sign
}
