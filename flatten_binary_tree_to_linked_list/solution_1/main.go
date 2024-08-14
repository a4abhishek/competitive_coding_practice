package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	// Unlink
	unlinked := root.Right // check for nil

	root.Right = root.Left
	root.Left = nil

	// Find the leaf
	leaf := root
	for leaf.Right != nil {
		leaf = leaf.Right
	}

	// Link
	leaf.Right = unlinked
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	flatten(root)
}
