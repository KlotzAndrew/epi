package lists_test

import (
	"reflect"
	"testing"

	"github.com/klotzandrew/epi/lists"
)

func TestMergeSorted(t *testing.T) {
	for _, tt := range []struct {
		x, y, expected []int
	}{
		{[]int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}},
		{[]int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{3}, []int{1, 2, 4}, []int{1, 2, 3, 4}},
	} {
		left := lists.NewList(tt.x)
		right := lists.NewList(tt.y)

		result := lists.MergeSorted(left, right)
		resultValues := result.ToSlice()
		if !reflect.DeepEqual(resultValues, tt.expected) {
			t.Errorf("MergeSorted(%v, %v) = %v, %v", tt.x, tt.y, resultValues, tt.expected)
		}
	}
}
