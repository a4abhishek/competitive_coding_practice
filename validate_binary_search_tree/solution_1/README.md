# Approach 1: Recursive Bound Validation

In this approach, we validate the binary search tree (BST) by checking each node's value against an allowable range (lower and upper bounds). This ensures that all nodes respect the BST property: left nodes are smaller than the current node, and right nodes are larger.

## Steps

1. **Recursion with Bounds**:
   - Traverse the tree using depth-first search (DFS).
   - For each node, check if the current value is within the allowable bounds.
   - For the left subtree, update the upper bound to be the current node’s value.
   - For the right subtree, update the lower bound to be the current node’s value.

2. **Base Case**:
   - If a node is `None`, return `true` (since an empty tree is a valid BST).

3. **Check Conditions**:
   - For each node, ensure its value is within the bounds (`lower_bound < node_value < upper_bound`).

### Example Walkthrough

Consider the following tree:

```python
    5
   / \
  1   4
     / \
    3   6
```

- Start with the root node `5` with initial bounds (`-∞, ∞`).
- Move to the left child `1`, with bounds updated to (`-∞, 5`). This node is valid since `1 < 5`.
- Move to the right child `4`, with bounds updated to (`5, ∞`). However, `4` is not greater than `5`, so the tree is not a valid BST.

### Why This Approach Works

By recursively updating the bounds as we move through the tree, we ensure that every node is checked against the values of its ancestors. This approach guarantees that the binary search tree property is maintained throughout the tree.

## Code Implementation

```python
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def is_valid_bst(root):
    def validate(node, low=-float('inf'), high=float('inf')):
        if not node:
            return True
        if not (low < node.val < high):
            return False
        return (validate(node.left, low, node.val) and
                validate(node.right, node.val, high))
    
    return validate(root)
```

### Complexity Analysis

- **Time Complexity:** `O(n)` where `n` is the number of nodes in the tree. We traverse every node once.
- **Space Complexity:** `O(h)` where `h` is the height of the tree. This is the space required by the recursion stack.
