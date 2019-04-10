package arrays_test

import (
	"reflect"
	"testing"

	"github.com/klotzandrew/epi/arrays"
)

func TestDuctFlagPartition(t *testing.T) {
	for _, tt := range []struct {
		pivot    int
		colors   []int
		expected []int
	}{
		{0, []int{1, 2, 0, 1, 0}, []int{0, 0, 1, 1, 2}},
		{2, []int{2, 2, 1, 1, 0}, []int{0, 1, 1, 2, 2}},
		{3, []int{2, 2, 1, 1, 0}, []int{0, 1, 1, 2, 2}},
	} {
		result := arrays.DutchFlagPartition(tt.pivot, tt.colors)
		if reflect.DeepEqual(result, tt.expected) != true {
			t.Errorf("DutchFlagPartition(%v) = %v, expected %v", tt.colors, result, tt.expected)
		}
	}
}
