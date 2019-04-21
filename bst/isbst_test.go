package bst_test

import (
	"testing"

	"github.com/klotzandrew/epi/bst"
)

func TestIsBST(t *testing.T) {
	for _, tt := range []struct {
		tree     *bst.BST
		expected bool
	}{
		{bst.NewBST(1, nil, nil), true},
		{bst.NewBST(
			1,
			bst.NewBST(2, nil, nil),
			nil,
		), false},
		{bst.NewBST(
			10,
			bst.NewBST(5, bst.NewBST(3, nil, nil), bst.NewBST(8, nil, nil)),
			bst.NewBST(15, nil, nil),
		), true},
	} {
		if result := bst.IsBST(tt.tree); result != tt.expected {
			t.Errorf("IsBST(%v) = %v, expected %v", tt.tree, result, tt.expected)
		}
	}
}
