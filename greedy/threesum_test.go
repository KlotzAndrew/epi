package greedy_test

import (
	"testing"

	"github.com/klotzandrew/epi/greedy"
)

func TestHasThreeSum(t *testing.T) {
	for _, tt := range []struct {
		arr []int
		sum int
		has bool
	}{
		{[]int{2, 3, 5, 7, 11}, 21, true},
	} {
		if result := greedy.HasThreeSum(tt.arr, tt.sum); result != tt.has {
			t.Errorf("HasThreeSum(%v, %d) = %v, expected %v", tt.arr, tt.sum, result, tt.has)
		}
	}
}
