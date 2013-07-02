package main

type PointSet struct {
	width int
	tree *avlTreeNode
}

func NewPointSet(width int) *PointSet {
	ps := new(PointSet)
	ps.width = width
	return ps
}

type avlTreeNode struct {
	left *avlTreeNode
	right *avlTreeNode
	val int
	h int
}

func makeAvlNode(n int) *avlTreeNode {
	a := new(avlTreeNode)
	a.val = n
	a.h = 1
	return a
}

func (ps *PointSet) Add(p Point) {
	pval := (p.Y * ps.width) + p.X + 1
	if ps.tree == nil {
		ps.tree = makeAvlNode(pval)
		return
	}
	ps.tree.insert(pval)
	checkAVL(&ps.tree)
}

func (ps *PointSet) Find(p Point) bool {
	pval := (p.Y * ps.width) + p.X + 1
	return ps.tree.find(pval)
}

func (n *avlTreeNode) find(val int) bool {
	if n == nil {
		return false
	}
	if n.val == val {
		return true
	}
	if val < n.val {
		return n.left.find(val)
	} else {
		return n.right.find(val)
	}
}

func (n *avlTreeNode) insert(val int) {
	if val < n.val {
		if n.left == nil {
			n.left = makeAvlNode(val)
		} else {
			n.left.insert(val)
			checkAVL(&n.left)
		}
	} else {
		if n.right == nil {
			n.right = makeAvlNode(val)
		} else {
			n.right.insert(val)
			checkAVL(&n.right)
		}
	}
	n.h = calcHeight(n)
}

func (n *avlTreeNode) height() int {
	if n == nil {
		return 0
	}
	return n.h
}

func checkAVL(n **avlTreeNode) {
	if *n == nil {
		return
	}
	bal := (*n).left.height() - (*n).right.height()
	if bal < 2 && bal > -2 {
		return //because its balanced yo
	}
	if bal > 1 {
		rotateRight(n)
	} else {
		rotateLeft(n)
	}
}

//If the height of the right subtree is more than one greater than the length of the left subtree
//then we need to balance it by rotating the node to the left
func rotateLeft(n **avlTreeNode) { //This node will be referred to as A
	A := *n

	//Move A's right child (Now referred to as B) up to its location
	*n = (*n).right

	//now set A's right child to B's left
	A.right = (*n).left

	//Put A on B's left
	(*n).left = cur

	//Update cached height values
	cur.h = calcHeight(cur)
	(*n).h = calcHeight(*n)
}

//See rotateLeft, same concept, just reversed
func rotateRight(n **avlTreeNode) {
	cur := *n
	*n =(*n).left
	cur.left = (*n).right
	(*n).right = cur
	cur.h = calcHeight(cur)
	(*n).h = calcHeight(*n)
}

func calcHeight(n *avlTreeNode) int {
	if n == nil {
		return 0
	} else {
		left := n.left.height()
		right := n.right.height()
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

func recurHeight(n *avlTreeNode) int {
	if n == nil {
		return 0
	}
	left := recurHeight(n.left)
	right := recurHeight(n.right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
