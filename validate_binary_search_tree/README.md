# [LeetCode 98: Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)

## Problem Statement

Given the `root` of a binary tree, determine if it is a valid binary search tree (BST).

### A valid BST is defined as follows

1. The left subtree of a node contains only nodes with keys **less than** the node's key.
2. The right subtree of a node contains only nodes with keys **greater than** the node's key.
3. Both the left and right subtrees must also be binary search trees.

## Examples

### Example 1

**Input:**

```python
    2
   / \
  1   3
```

**Output:** `true`

### Example 2

**Input:**

```python
    5
   / \
  1   4
     / \
    3   6
```

**Output:** `false`

**Explanation:**

- The root node's value is 5, but its right child's value is 4, which violates the binary search tree property.

## Constraints

- The number of nodes in the tree is in the range `[1, 10^4]`.
- `-2^31 <= Node.val <= 2^31 - 1`

## Approaches

- [Approach 1: Recursive Bound Validation](./solution_1/)
- [Approach 2: In-order Traversal with Generator](./solution_2/)

## Common Pitfalls

### Incorrect Approach: Simple Child Comparison

In this incorrect approach, the validation only checks whether each node's left child is smaller and the right child is larger. However, this method fails to validate the full tree because it doesn't check for the relationship with ancestor nodes.

```python
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def isValidBST(root):
    if root is None:
        return True

    if root.left and root.left.val >= root.val:
        return False

    if root.right and root.right.val <= root.val:
        return False

    return isValidBST(root.left) and isValidBST(root.right)
```

### Why This Fails

The problem with this approach is that it only compares the immediate children with their parent, but it doesn't ensure that all nodes in the left subtree are less than the current node and all nodes in the right subtree are greater than the current node. For example, in the following tree:

```python
    10
   /  \
  5    15
      /  \
     6    20
```

The node `6` is in the right subtree of `10`, so it should be greater than `10`, but the simple comparison only checks that `6 < 15`, which is incorrect. As a result, this approach will return `True` for an invalid BST.
