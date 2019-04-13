package lists

type Node struct {
	value int
	next  *Node
}

func NewList(values []int) *Node {
	head := &Node{value: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		n := Node{value: values[i]}
		current.next = &n
		current = current.next
	}

	return head
}

func (n *Node) ToSlice() (s []int) {
	for c := n; c != nil; c = c.next {
		s = append(s, c.value)
	}
	return
}

func MergeSorted(x, y *Node) *Node {
	dummy := &Node{}
	current := dummy

	left, right := x, y
	for left != nil || right != nil {
		switch {
		case left == nil:
			current.next = right
			right = right.next
		case right == nil:
			current.next = left
			left = left.next
		case left.value <= right.value:
			current.next = left
			left = left.next
		case left.value > right.value:
			current.next = right
			right = right.next
		}
		current = current.next
	}

	return dummy.next
}
