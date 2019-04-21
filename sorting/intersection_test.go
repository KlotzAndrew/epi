package sorting_test

import (
	"reflect"
	"testing"

	"github.com/klotzandrew/epi/sorting"
)

func TestIntersect(t *testing.T) {
	for _, tt := range []struct {
		a, b, expected []int
	}{
		{[]int{1, 3, 5}, []int{1, 2, 3, 4}, []int{1, 3}},
		{[]int{1, 3, 5}, []int{1, 2, 3, 3, 4}, []int{1, 3}},
	} {
		if result := sorting.Intersect(tt.a, tt.b); !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Intersect(%v, %v) = %v, expected %v", tt.a, tt.b, result, tt.expected)
		}
	}
}
