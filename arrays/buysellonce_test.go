package arrays_test

import (
	"testing"

	"github.com/klotzandrew/epi/arrays"
)

func TestMaxProfitOnce(t *testing.T) {
	for _, tt := range []struct {
		prices []int
		max    int
	}{
		{[]int{10, 5, 20}, 15},
		{[]int{1, 5, 20}, 19},
	} {
		if result := arrays.MaxProfitOnce(tt.prices); result != tt.max {
			t.Errorf("MaxProfitOnce(%d) = %d, %d", tt.prices, result, tt.max)
		}
	}
}
