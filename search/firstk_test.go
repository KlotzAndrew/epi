package search_test

import (
	"testing"

	"github.com/klotzandrew/epi/search"
)

func TestFirstK(t *testing.T) {
	for _, tt := range []struct {
		arr              []int
		target, expected int
	}{
		{[]int{1, 2, 3, 4}, 3, 2},
		{[]int{1, 2, 3, 4}, 33, -1},
		{[]int{-10, -5, 0, 1, 10}, 0, 2},
	} {
		if result := search.FirstK(tt.arr, tt.target); result != tt.expected {
			t.Errorf("FirstK(%v) = %v, expected %v", tt.arr, result, tt.expected)
		}
	}
}
