# [LeetCode - 108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/)

**Difficulty:** Easy

## Problem Statement

Given an integer array `nums` where the elements are sorted in **ascending order**, convert it to a <span title="A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.">**height-balanced** binary search tree.

## Examples

### Example 1:

**Input:** nums = `[-10,-3,0,5,9]`

**Output:**
```
        0
       / \
     -3   9
     /   /
  -10   5
```

**Explanation:** `[0,-3,9,-10,null,5]` is also accepted.

Another valid output could be:
```
        0
       / \
    -10   5
      \    \
     -3    9
```

### Example 2:

**Input:** nums = `[1,3]`

**Output:**
```
  3       1
 /         \
1           3
```

**Explanation:** `[1,null,3]` and `[3,1]` are both height-balanced BSTs.

## Constraints:

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` is sorted in a **strictly increasing order**.
