package bst

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

type BST struct {
	value       int
	left, right *BST
}

func NewBST(v int, left, right *BST) *BST {
	return &BST{v, left, right}
}

func IsBST(tree *BST) bool {
	return areKeysInRange(tree, minInt, maxInt)
}

func areKeysInRange(t *BST, min, max int) bool {
	if t == nil {
		return true
	}

	if t.value < min || max < t.value {
		return false
	}

	return areKeysInRange(t.left, min, t.value) && areKeysInRange(t.right, t.value, max)
}
