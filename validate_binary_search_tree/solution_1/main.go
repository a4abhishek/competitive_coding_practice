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

func validate(low int, root *TreeNode, high int) bool {
	if root == nil {
		return true
	}

	if low >= root.Val || root.Val >= high {
		return false
	}

	return validate(low, root.Left, root.Val) && validate(root.Val, root.Right, high)
}

func isValidBST(root *TreeNode) bool {
	return validate(math.MinInt, root, math.MaxInt)
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
