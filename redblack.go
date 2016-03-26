// main implements a simple red/black tree of ints
//
// Built following:
// http://www.eternallyconfuzzled.com/tuts/datastructures/jsw_tut_rbtree.aspx
package main

import (
	"fmt"
)

// RBNode represents a node in the red/black tree.
type RBNode struct {
	Val         int
	Red         bool
	Left, Right *RBNode
}

// Insert inserts creates a new red node and inserts it into the tree, it then
// sets the root as black and returns it.
func (node *RBNode) Insert(val int) *RBNode {
	node = node.insertRecursive(val)
	node.Red = false
	return node
}

func (node *RBNode) insertRecursive(val int) *RBNode {
	if node == nil {
		return &RBNode{val, true, nil, nil}
	}
	// Only insert if not a duplicate
	if val != node.Val {
		if val < node.Val {
			node.Left = node.Left.insertRecursive(val)
			if node.Left.IsRed() {
				if node.Right.IsRed() {
					// Case 1: Both children red
					node.Red = true
					node.Left.Red = false
					node.Right.Red = false
				} else {
					if node.Left.Left.IsRed() {
						// Case 2: Red violation on left child
						node = node.Rot1Right()
					} else if node.Left.Right.IsRed() {
						// Case 3: Red violation on right child
						node = node.Rot2Right()
					}
				}
			}
		} else {
			node.Right = node.Right.insertRecursive(val)
			if node.Right.IsRed() {
				if node.Left.IsRed() {
					// Case 1: Both children red
					node.Red = true
					node.Left.Red = false
					node.Right.Red = false
				} else {
					if node.Right.Right.IsRed() {
						// Case 2: Red violation on right child
						node = node.Rot1Left()
					} else if node.Right.Left.IsRed() {
						// Case 3: Red violation on left child
						node = node.Rot2Left()
					}
				}
			}
		}

	}
	return node
}

// IsRed returns true if this node is red, false if it is black or nil
func (node *RBNode) IsRed() bool {
	if node == nil {
		return false
	}
	return node.Red
}

// Rot1Left rotates this node to the left, returns the node to replace it with
// (its right child)
func (node *RBNode) Rot1Left() *RBNode {
	// Rotate
	save := node.Right
	node.Right = save.Left
	save.Left = node

	// Recolor
	node.Red = true
	save.Red = false
	return save
}

// Rot2Left rotates this nodes right child to the right, then rotates this node
// to the left, returning the node to replace it with (its new right child)
func (node *RBNode) Rot2Left() *RBNode {
	// Rotate right child to the right
	node.Right = node.Right.Rot1Right()
	fmt.Println(node)
	// Then rotate this node left
	x := node.Rot1Left()
	fmt.Println(x)
	return x
}

// Rot1Right rotates this node to the right, returns the node to replace it with
// (its left child)
func (node *RBNode) Rot1Right() *RBNode {
	// Rotate
	save := node.Left
	node.Left = save.Right
	save.Right = node

	// Recolor
	node.Red = true
	save.Red = false
	return save
}

// Rot2Right rotates this nodes left child to the left, then rotates this node
// to the right, returning the node to replace it with (its new left child)
func (node *RBNode) Rot2Right() *RBNode {
	// Rotate left child to the left
	node.Left = node.Left.Rot1Left()
	fmt.Println(node)
	// Then rotate this node right
	x := node.Rot1Right()
	fmt.Println(x)
	return x
}

// Validate tests the tree below the given node to make sure it does not
// contain any red or black violations, is a correct binary tree, and has a
// consistent black height.  Returns black height or an error.  This function
// is slow and recursive.
func (node *RBNode) Validate() (int, error) {
	if node == nil {
		// Consider nil leafs black
		return 1, nil
	}
	left, right := node.Left, node.Right

	if node.IsRed() {
		if left.IsRed() || right.IsRed() {
			return 0, fmt.Errorf("Red violation at value %v", node.Val)
		}
	}

	if (left != nil && node.Val <= left.Val) ||
		(right != nil && right.Val <= node.Val) {
		return 0, fmt.Errorf("Binary tree violation at value %v", node.Val)
	}

	lheight, err := left.Validate()
	if err != nil {
		return 0, err
	}
	rheight, err := right.Validate()
	if err != nil {
		return 0, err
	}
	if lheight != rheight {
		return 0, fmt.Errorf("Black height violation at value %v", node.Val)
	}

	// Only count black nodes
	if node.IsRed() {
		return lheight, nil
	}
	return lheight + 1, nil
}

func (node *RBNode) String() string {
	if node == nil {
		return ""
	}
	color := "B"
	if node.Red {
		color = "R"
	}
	return fmt.Sprintf("%s:%v(%s,%s)", color, node.Val, node.Left, node.Right)
}
