package main

import (
	"testing"
)

// Helper function to check if a tree is a BST
func isBST(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
		return false
	}

	return isBST(node.Left, min, &node.Val) && isBST(node.Right, &node.Val, max)
}

// Helper function to check if a tree is height-balanced
func isBalanced(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	leftBalanced, leftHeight := isBalanced(node.Left)
	rightBalanced, rightHeight := isBalanced(node.Right)

	if !leftBalanced || !rightBalanced {
		return false, 0
	}

	if abs(leftHeight-rightHeight) > 1 {
		return false, 0
	}

	return true, max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Test function for sortedArrayToBST
func TestSortedArrayToBST(t *testing.T) {
	tests := [][]int{
		{-10, -3, 0, 5, 9},
		{1, 3},
		{},
		{1},
		{1, 2, 3, 4, 5, 6, 7},
	}

	for _, test := range tests {
		tree := sortedArrayToBST(test)

		// Check if the tree is a valid BST
		if !isBST(tree, nil, nil) {
			t.Errorf("Tree from input %v is not a valid BST", test)
		}

		// Check if the tree is height-balanced
		balanced, _ := isBalanced(tree)
		if !balanced {
			t.Errorf("Tree from input %v is not height-balanced", test)
		}
	}
}
