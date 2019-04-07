package ptypes

var reverseBitTable [1 << 16]uint16

func init() {
	for i := 0; i < len(reverseBitTable); i++ {
		reverseBitTable[i] = uint16(reverseBits(uint64(i)) >> 48)
	}
}

func reverseBits(x uint64) uint64 {
	for i := uint64(0); i < 64/2; i++ {
		x = SwapBits(x, i, 63-i)
	}
	return x
}

func ReverseBits(x uint64) uint64 {
	const bitMask uint64 = 0xffff
	const maskSize uint = 16

	return uint64(reverseBitTable[(x>>(0*maskSize))&bitMask])<<(3*maskSize) |
		uint64(reverseBitTable[(x>>(1*maskSize))&bitMask])<<(2*maskSize) |
		uint64(reverseBitTable[(x>>(2*maskSize))&bitMask])<<(1*maskSize) |
		uint64(reverseBitTable[(x>>(3*maskSize))&bitMask])<<(0*maskSize)
}
