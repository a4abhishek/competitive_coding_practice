# [LeetCode 2385: Amount of Time for Binary Tree to be Burnt](https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/)

## Problem Statement

You are given a binary tree and a specific node, referred to as the *target*. Your task is to determine the minimum time required to burn the entire binary tree if the target node is set on fire. In this problem, it is known that all nodes connected to the target (i.e., its left child, right child, and parent) will also catch fire after 1 second.

### Notes

- The tree contains unique values.
- The fire spreads simultaneously to all connected nodes every second.

## Examples

### Example 1

**Input:**

```
          1
        /   \
      2      3
    /  \      \
   4    5      6
       / \      \
      7   8      9
                   \
                   10
Target Node = 8
```

**Output:** `7`

**Explanation:**

- If the leaf node with the value 8 is set on fire:
  - After 1 second: Node 5 is set on fire.
  - After 2 seconds: Nodes 2 and 7 are set on fire.
  - After 3 seconds: Nodes 4 and 1 are set on fire.
  - After 4 seconds: Node 3 is set on fire.
  - After 5 seconds: Node 6 is set on fire.
  - After 6 seconds: Node 9 is set on fire.
  - After 7 seconds: Node 10 is set on fire.
- It takes 7 seconds to burn the complete tree.

### Example 2

**Input:**

```
          1
        /   \
      2      3
    /  \      \
   4    5      7
  /    / 
 8    10
Target Node = 10
```

**Output:** `5`

**Explanation:**

- If node 10 is set on fire:
  - After 1 second: Node 5 is set on fire.
  - After 2 seconds: Node 2 is set on fire.
  - After 3 seconds: Node 1 is set on fire.
  - After 4 seconds: Node 3 is set on fire.
  - After 5 seconds: Node 7 is set on fire.
- It takes 5 seconds to burn the complete tree.

## Constraints

- `1 ≤ number of nodes ≤ 10^5`
- `1 ≤ values of nodes ≤ 10^5`

## Approach

To solve the **Burning Tree** problem, we can follow these steps:

### Steps

1. **Tree Traversal & Parent Tracking:**
   - First, traverse the binary tree to establish the parent-child relationships. This is important because while the fire can spread to the left and right children, it can also spread upwards to the parent.

2. **Multi-source BFS from the Target:**
   - Once the target node is on fire, all nodes connected to it will catch fire simultaneously. This forms a multi-level propagation, and a Breadth-First Search (BFS) is ideal for simulating the fire spread.
   - We can start BFS from the target node and simulate the spread of fire, visiting its neighbors (left child, right child, and parent) in each step.

3. **Minimizing the Time:**
   - Track the time it takes to burn each node. The process stops when all nodes are burned, and the maximum time value from the BFS simulation will be our result.

### Detailed Explanation

1. Use a DFS to traverse the tree and store each node's parent. This way, for any node, we know both its children and parent.
2. Perform a BFS starting from the target node. Initially, mark the target node as burned and push it into the queue. For each node in the queue, propagate the fire to its unburned neighbors (left child, right child, parent).
3. Keep a counter to track the time taken to burn all nodes.
4. The BFS will spread the fire level by level, and once all nodes are processed, the maximum value of the counter will be the minimum time required to burn the entire tree.

### Example Walkthrough

Consider the tree structure from **Example 1**:

- **Initial State**:
  - Target node = 8
  - BFS starts at node 8, marking node 5 (its parent) as next to burn.

- **After 1 Second**:
  - Node 5 is on fire.
  - Nodes 7 (left child of 5) and 2 (parent of 5) will burn next.

- **After 2 Seconds**:
  - Nodes 7 and 2 are on fire.
  - Node 4 (left child of 2) and 1 (parent of 2) will burn next.

- **After 3 Seconds**:
  - Nodes 4 and 1 are on fire.
  - Node 3 (right child of 1) will burn next.

- **After 4 Seconds**:
  - Node 3 is on fire.
  - Node 6 (right child of 3) will burn next.

- **After 5 Seconds**:
  - Node 6 is on fire.
  - Node 9 (right child of 6) will burn next.

- **After 6 Seconds**:
  - Node 9 is on fire.
  - Node 10 (right child of 9) will burn next.

- **After 7 Seconds**:
  - Node 10 is on fire. The entire tree is burned.

## Example Implementation (Python)

```python
class TreeNode:
    def __init__(self, data=0, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right

def min_time_to_burn_tree(root, target):
    parent_map = {}
    
    def dfs(node, parent=None):
        if not node:
            return
        parent_map[node] = parent
        dfs(node.left, node)
        dfs(node.right, node)
    
    dfs(root)
    
    # BFS to simulate burning
    from collections import deque
    queue = deque()
    visited = set()
    time = -1
    
    # Find target node
    def find_target(node):
        if not node:
            return None
        if node.data == target:
            return node
        return find_target(node.left) or find_target(node.right)
    
    target_node = find_target(root)
    queue.append(target_node)
    visited.add(target_node)
    
    while queue:
        time += 1
        for _ in range(len(queue)):
            node = queue.popleft()
            
            # Check left child, right child, and parent
            for neighbor in (node.left, node.right, parent_map[node]):
                if neighbor and neighbor not in visited:
                    visited.add(neighbor)
                    queue.append(neighbor)
    
    return time
```

## Complexity Analysis

- **Time Complexity:** `O(N)` where `N` is the number of nodes in the tree. We traverse each node once in both the DFS and BFS.
- **Space Complexity:** `O(H)` where `H` is the height of the tree, due to the space used by the recursion stack during DFS and the BFS queue.
