package btree

type BTree struct {
	value               int
	parent, left, right *BTree
}

func NewTree(values []int) *BTree {
	head := &BTree{value: values[0]}

	for i := 1; i < len(values); i++ {
		newNode := &BTree{value: values[i]}
		head.addNode(newNode)
	}

	return head
}

func (parent *BTree) IsBalanced() bool {
	if parent == nil {
		return true
	}

	lh := parent.left.height()
	rh := parent.right.height()
	diff := lh - rh
	if (diff >= 0 && diff <= 1) || (diff < 0 && diff >= -1) {
		lb := parent.left.IsBalanced()
		rb := parent.right.IsBalanced()
		return lb && rb
	}
	return false
}

func (parent *BTree) height() int {
	if parent == nil {
		return 0
	}
	hl := 1 + parent.left.height()
	hr := 1 + parent.right.height()

	if hl > hr {
		return hl
	}
	return hr
}

func (parent *BTree) addNode(new *BTree) {
	if new.value < parent.value {
		if parent.left == nil {
			parent.left = new
			new.parent = parent
			return
		}
		parent.left.addNode(new)
	} else {
		if parent.right == nil {
			parent.right = new
			new.parent = parent
			return
		}
		parent.right.addNode(new)
	}
}
