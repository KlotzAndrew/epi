package heaps_test

import (
	"reflect"
	"testing"

	"github.com/klotzandrew/epi/heaps"
)

func TestMergeSorted(t *testing.T) {
	for _, tt := range []struct {
		input    [][]int
		expected []int
	}{
		{[][]int{{1, 3}, {2, 4}}, []int{1, 2, 3, 4}},
	} {
		if result := heaps.MergeSorted(tt.input); !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("MergeSorted(%v) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}
