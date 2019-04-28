package ptypes

// Power .
func Power(x, y uint64) uint64 {
	var result uint64 = 1
	power := y

	for power != 0 {
		if (power & 1) != 0 {
			result *= x
		}
		x *= x
		power >>= 1
	}
	return result
}
