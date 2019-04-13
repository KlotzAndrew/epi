package stacks_test

import (
	"testing"

	"github.com/klotzandrew/epi/stacks"
)

func TestMax(t *testing.T) {
	for _, tt := range []struct {
		values []int
		pops   int
		max    int
	}{
		{[]int{2, 3, 4, 3}, 0, 4},
		{[]int{2, 3, 4, 3}, 2, 3},
	} {
		s := stacks.NewIntStack()
		for _, v := range tt.values {
			s.Push(v)
		}
		for i := 0; i < tt.pops; i++ {
			s.Pop()
		}
		if result := s.Max(); result != tt.max {
			t.Errorf("Max %d = %d, %d", tt.values, result, tt.max)
		}
	}
}
