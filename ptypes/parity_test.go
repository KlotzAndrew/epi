package ptypes_test

import (
	"testing"

	"github.com/klotzandrew/epi/ptypes"
)

const (
	odd  = 1
	even = 0
)

func TestParity(t *testing.T) {
	for _, tt := range []struct {
		value    uint64
		expected uint8
	}{
		{0, even},
		{1, odd},
		{2, odd},
		{3, even},
		{4, odd},
		{5, even},
		{6, even},
		{7, odd},
		{8, odd},
		{9, even},
		{998, odd},
		{999, even},
	} {
		if result := ptypes.Parity(tt.value); result != tt.expected {
			t.Errorf("%b equals %d; want %d", tt.value, result, tt.expected)
		}
	}
}
