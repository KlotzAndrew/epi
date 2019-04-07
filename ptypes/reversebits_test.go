package ptypes_test

import (
	"testing"

	"github.com/klotzandrew/epi/ptypes"
)

func TestReverseBits(t *testing.T) {
	for _, tt := range []struct {
		value    uint64
		expected uint64
	}{
		{1, 1 << 63},
		{2, 1 << 62},
		{4, 1 << 61},
		{8, 1 << 60},
		{0x00000000ffffffff, 0xffffffff00000000},
		{0x00000000000000ff, 0xff00000000000000},
		{0x00000000ff000000, 0x000000ff00000000},
		{0xffffffffffffffff, 0xffffffffffffffff},
	} {
		if result := ptypes.ReverseBits(tt.value); result != tt.expected {
			t.Errorf("ReverseBits(%b) = %b, expected %b", tt.value, result, tt.expected)
		}
	}
}
