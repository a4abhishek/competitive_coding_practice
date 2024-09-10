# Approach 2: In-order Traversal with Generator

## Approach Overview

This approach leverages the properties of an **in-order traversal** for binary search trees. The in-order traversal visits nodes in increasing order if the tree is a valid BST. The idea is to traverse the tree in in-order fashion and ensure that the values appear in strictly increasing order by comparing the current node value with the previous node value.

### Key Idea

- In a **Binary Search Tree (BST)**, for each node, all values in its left subtree must be **less** than the node’s value, and all values in its right subtree must be **greater** than the node’s value.
- An **in-order traversal** of a BST visits nodes in increasing order.
- By comparing consecutive node values during the in-order traversal, we can determine whether the tree is a valid BST.

### Plan

1. **In-order Traversal**:
   - We perform an in-order traversal of the tree, which visits the left child, then the current node, and finally the right child.
   - This traversal is implemented using a Python **generator** to yield node values one by one.
2. **Strictly Increasing Check**:
   - As we traverse the tree, we compare the value of each node with the value of the previously visited node. If the current node’s value is less than or equal to the previous node’s value, the tree is not a valid BST.

## Code Implementation

```python
# Define the TreeNode structure
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# In-order generator function
def inOrderNode(root):
    if root is None:
        return
    # Recursively yield from the left subtree
    yield from inOrderNode(root.left)
    # Yield the current node's value
    yield root.val
    # Recursively yield from the right subtree
    yield from inOrderNode(root.right)

# Function to check if the tree is a valid BST
def isValidBST(root):
    inOrder = inOrderNode(root)  # Get the in-order sequence
    try:
        prevNode = next(inOrder)  # Get the first node value
    except StopIteration:
        return True  # An empty tree is a valid BST
    
    for node in inOrder:
        if node > prevNode:
            prevNode = node  # Update previous node
        else:
            return False  # If the current node is not greater, it's not a valid BST
    return True
```

### Explanation of the Code

1. **`inOrderNode` function**: This is a recursive generator function that performs an in-order traversal of the tree, yielding each node’s value.
   - If the node is `None`, it returns immediately.
   - It yields all values from the left subtree, then the current node’s value, and finally all values from the right subtree.
2. **`isValidBST` function**: This function uses the `inOrderNode` generator to iterate over the tree values in in-order and checks whether the values follow a strictly increasing order.
   - The first node's value is captured in `prevNode`.
   - Then, for each subsequent node value, it checks if the current node's value is greater than `prevNode`. If not, the function returns `False` indicating the tree is not a valid BST.
   - If the traversal completes without violations, the function returns `True`, confirming the tree is a valid BST.

## Example Walkthrough

### Tree Example 1

```python
    2
   / \
  1   3
```

For this tree, the in-order traversal will yield the values `[1, 2, 3]`, which are in strictly increasing order. Therefore, the tree is a valid BST.

### Tree Example 2

```python
    5
   / \
  1   4
     / \
    3   6
```

For this tree, the in-order traversal will yield the values `[1, 5, 3, 4, 6]`. Since the value `3` comes after `5`, which violates the BST property, the tree is **not** a valid BST.

## Time and Space Complexity

### Time Complexity

- **O(n)**, where `n` is the number of nodes in the tree. This is because we visit each node exactly once during the in-order traversal.

### Space Complexity

- **O(h)**, where `h` is the height of the tree. This is the space required for the recursive call stack during the in-order traversal. In the worst case, `h` can be as large as `n` for a completely unbalanced tree.

## Advantages of This Approach

1. **Efficient**: Since we stop as soon as we find a violation of the BST property, the traversal may terminate early in some cases.
2. **Simple and Elegant**: Using an in-order traversal ensures that we naturally check the BST properties without needing extra data structures.
3. **Lazy Evaluation**: The use of a generator allows for lazy evaluation of the nodes, which can help in reducing memory usage and improving performance in large trees.
