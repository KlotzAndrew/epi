package ptypes

func Multiply(x, y uint64) uint64 {
	var sum uint64 = 0
	for x != 0 {
		if (x & 1) != 0 {
			sum = add(sum, y)
		}
		x >>= 1
		y <<= 1
	}
	return sum
}

func add(x, y uint64) uint64 {
	for y != 0 {
		var carry = x & y
		x = x ^ y
		y = carry << 1
	}
	return x
}
