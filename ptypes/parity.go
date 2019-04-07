package ptypes

func Parity(x uint64) uint8 {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return uint8(x & 1)
}
