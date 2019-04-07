package ptypes_test

import (
	"testing"

	"github.com/klotzandrew/epi/ptypes"
)

func TestClosestInt(t *testing.T) {
	for _, tt := range []struct {
		value    uint64
		expected uint64
	}{
		{0, 0},
		{1, 2},
		{2, 1},
		{4, 2},
		{5, 6},
		{7, 11},
		{8, 4},
		{9, 10},
	} {
		if result, _ := ptypes.ClosestInt(tt.value); result != tt.expected {
			t.Errorf("ClosestInt(%b) = %b, expected %b", tt.value, result, tt.expected)
		}
	}
}
