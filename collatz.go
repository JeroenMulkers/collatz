package collatz

var (
	Tree = &Node{1, nil, nil, nil, 0}
	End  = &Node{0, nil, nil, nil, 0}
	Lut  = make(map[int]*Node)
)

func init() {
	Lut[0] = End
	Lut[1] = Tree
}

type Node struct {
	value int
	down  *Node
	left  *Node
	right *Node
	order int
}

func NewNode(value int) *Node {
	n := &Node{value, nil, nil, nil, -1}
	Lut[value] = n
	if n.value != 0 {
		n.left, _ = Lut[n.leftValue()]
		n.right, _ = Lut[n.rightValue()]
		n.down, _ = Lut[n.downValue()]
	}
	if n.down != nil {
		n.order = n.down.Order() + 1
	}
	return n
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
		return 0
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
		n.down = NewNode(n.downValue())
	}
	return n.down
}

func (n *Node) Left() *Node {
	if n.left == nil {
		n.left = NewNode(n.leftValue())
	}
	return n.left
}

func (n *Node) Right() *Node {
	if n.right == nil {
		n.right = NewNode(n.rightValue())
	}
	return n.right
}
