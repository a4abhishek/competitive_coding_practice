package main

import (
	"fmt"
	"math"
)

// Tree DataStructure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

// prevVal is the last value before the tree starting this node
func inOrderValidation(prevVal *int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !inOrderValidation(prevVal, root.Left) {
		return false
	}

	if *prevVal >= root.Val {
		return false
	}

	*prevVal = root.Val

	return inOrderValidation(prevVal, root.Right)
}

func isValidBST(root *TreeNode) bool {
	prevVal := math.MinInt
	return inOrderValidation(&prevVal, root)
}

func main() {
	var tree *TreeNode

	// Tree 1
	tree = NewNode(2)
	tree.Left = NewNode(1)
	tree.Right = NewNode(3)

	fmt.Println(isValidBST(tree))

	// Tree 2
	tree = NewNode(5)
	tree.Left = NewNode(1)
	tree.Right = NewNode(4)
	tree.Right.Left = NewNode(3)
	tree.Right.Right = NewNode(6)

	fmt.Println(isValidBST(tree))

	// Tree 3
	tree = NewNode(5)
	tree.Left = NewNode(4)
	tree.Right = NewNode(6)
	tree.Right.Left = NewNode(3)
	tree.Right.Right = NewNode(7)

	fmt.Println(isValidBST(tree))
}
