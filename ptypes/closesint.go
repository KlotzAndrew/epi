package ptypes

func ClosestInt(x uint64) (uint64, bool) {
	for i := uint64(0); i < 63; i++ {
		if (x>>i)&1 != (x>>(i+1))&1 {
			x ^= 1<<i | 1<<(i+1)
			return x, true
		}
	}
	return 0, false
}
