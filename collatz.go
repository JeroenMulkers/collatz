package collatz

/*
   Basic structure of the Collatz tree which will be build at the fly
*/

var (
	End = &Node{0, nil, nil, nil, 0}
	Lut = make(map[int]*Node)
)

func init() {
	Lut[0] = End
}

type Node struct {
	value int
	down  *Node
	left  *Node
	right *Node
	order int
}

// Create new node and attach to the existing tree
func NewNode(value int) *Node {
	if value < 1 {
		return End
	}
	n := &Node{value, nil, nil, nil, -1}
	// attach to the existing tree
	Lut[value] = n
	if n.value != 0 {
		n.left, _ = Lut[n.leftValue()]
		n.right, _ = Lut[n.rightValue()]
		n.down, _ = Lut[n.downValue()]
	}
	// add order only if order of down node is known
	if n.down != nil && n.down.order != -1 {
		n.order = n.down.order + 1
	}
	return n
}

// Return the value of the node
func (n *Node) Value() int {
	return n.value
}

// Return the value of the down node
func (n *Node) downValue() int {
	if n.value%2 == 0 {
		return n.value / 2
	} else {
		return 3*n.value + 1
	}
}

// Return the value of the left node
func (n *Node) leftValue() int {
	return n.value * 2
}

// Return the value of the down node
func (n *Node) rightValue() int {
	if ((n.value - 1) % 3) == 0 {
		return (n.value - 1) / 3
	} else {
		return 0
	}
}

// Return the order of the node, go down if necessary
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

// Return the node below (create if it does not exist)
func (n *Node) Down() *Node {
	if n.down == nil {
		n.down = NewNode(n.downValue())
	}
	return n.down
}

// Return the left node (create if it does not exist)
func (n *Node) Left() *Node {
	if n.left == nil {
		n.left = NewNode(n.leftValue())
	}
	return n.left
}

// Return the right node (create if it does not exist)
func (n *Node) Right() *Node {
	if n.right == nil {
		n.right = NewNode(n.rightValue())
	}
	return n.right
}
