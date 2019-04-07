package ptypes_test

import (
	"testing"

	"github.com/klotzandrew/epi/ptypes"
)

func TestSwapBits(t *testing.T) {
	for _, tt := range []struct {
		x, i, j  uint64
		expected uint64
	}{
		{3, 0, 1, 3},
		{5, 0, 1, 6},
		{11, 1, 2, 13},
	} {
		if result := ptypes.SwapBits(tt.x, tt.i, tt.j); result != tt.expected {
			t.Errorf("SwapBits(%b, %d, %d) = %b; want %b", tt.x, tt.i, tt.j, result, tt.expected)
		}
	}
}
