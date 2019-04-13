package btree_test

import (
	"testing"

	"github.com/klotzandrew/epi/btree"
)

func TestIsBalanced(t *testing.T) {
	for _, tt := range []struct {
		balanced bool
		values   []int
	}{
		{true, []int{3, 2, 4}},
		{false, []int{1, 2, 3, 4}},
		{false, []int{1, 1, 1, 1}},
	} {
		tree := btree.NewTree(tt.values)

		if result := tree.IsBalanced(); result != tt.balanced {
			t.Errorf("IsBalanced(%v) = %v, %v", tt.values, result, tt.balanced)
		}
	}
}
