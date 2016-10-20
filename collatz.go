package collatz

var (
	Tree = &Node{1, nil, nil, nil, 0}
)

type Node struct {
	value int
	down  *Node
	left  *Node
	right *Node
	order int
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil, nil, -1}
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) leftValue() int {
	return n.value * 2
}

func (n *Node) downValue() int {
	if n.value%2 == 0 {
		return n.value / 2
	} else {
		return 3*n.value + 1
	}
}

func (n *Node) rightValue() int {
	if ((n.value - 1) % 3) == 0 {
		return (n.value - 1) / 3
	} else {
		return -1
	}
}

func (n *Node) Order() int {
	if n.order == -1 {
		if n.value == 1 {
			n.order = 1
		} else { // Go down to determine order (~recursive)
			n.order = n.Down().Order() + 1
		}
	}
	return n.order
}

func (n *Node) Down() *Node {
	if n.down == nil {
		if n.value%2 == 0 {
			n.down = &Node{n.downValue(), nil, n, nil, -1}
		} else {
			n.down = &Node{n.downValue(), nil, nil, n, -1}
		}
	}
	return n.down
}

func (n *Node) Left() *Node {
	if n.left == nil {
		n.left = &Node{n.leftValue(), n, nil, nil, n.order + 1}
	}
	return n.left
}

func (n *Node) Right() *Node {
	if n.left == nil && ((n.value-1)%3) == 0 {
		n.right = &Node{n.rightValue(), n, nil, nil, n.order + 1}
	}
	return n.right
}
