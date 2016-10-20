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

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Order() int {
	return n.order
}

func (n *Node) Down() *Node {
	return n.down
}

func (n *Node) Left() *Node {
	if n.left == nil {
		n.left = &Node{2 * n.value, n, nil, nil, n.order + 1}
	}
	return n.left
}

func (n *Node) Right() *Node {
	if n.left == nil && ((n.value-1)%3) == 0 {
		n.right = &Node{(n.value - 1) / 3, n, nil, nil, n.order + 1}
	}
	return n.right
}
