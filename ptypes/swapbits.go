package ptypes

func SwapBits(x, i, j uint64) uint64 {
	if ((x >> i) & 1) != ((x >> j) & 1) {
		x ^= (1<<i | 1<<j)
	}
	return x
}
