package ptypes_test

import (
	"testing"

	"github.com/klotzandrew/epi/ptypes"
)

func TestMultiply(t *testing.T) {
	for _, tt := range []struct {
		x, y, expected uint64
	}{
		{2, 2, 4},
		{3, 3, 9},
		{12, 12, 144},
		{99, 10, 990},
	} {
		if result := ptypes.Multiply(tt.x, tt.y); result != tt.expected {
			t.Errorf("Multiply(%v, %v) = %v, expected %v", tt.x, tt.y, result, tt.expected)
		}
	}
}
