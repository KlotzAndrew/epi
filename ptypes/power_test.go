package ptypes_test

import (
	"testing"

	"github.com/klotzandrew/epi/ptypes"
)

func TestPower(t *testing.T) {
	for _, tt := range []struct {
		x, y, expected uint64
	}{
		{2, 2, 4},
		{2, 4, 16},
		{3, 3, 27},
	} {
		if result := ptypes.Power(tt.x, tt.y); result != tt.expected {
			t.Errorf("Power(%v, %v) = %v, exp %v", tt.x, tt.y, result, tt.expected)
		}
	}
}
