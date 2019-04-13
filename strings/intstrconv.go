package strings

func StringToInt(str string) int64 {
	result, neg := uint64(0), int64(1)

	for _, c := range str {
		if c == '-' {
			neg = -1
			continue
		}
		result = result*10 + uint64(c-'0')
	}

	return int64(result) * neg
}

func IntToString(v int64) string {
	var result [20]byte
	i := len(result)

	neg := false
	if v < 0 {
		neg = true
		v = -v
	}

	if v == 0 {
		return "0"
	}

	for v > 0 {
		i--
		result[i] = byte(v%10 + '0')
		v /= 10
	}

	if neg {
		i--
		result[i] = '-'
	}

	return string(result[i:])
}
