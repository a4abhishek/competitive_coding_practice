# [LeetCode 4: Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## Problem Statement

You are given two sorted arrays `nums1` of size `m` and `nums2` of size `n`. Your task is to find the **median** of the two sorted arrays combined.

The overall time complexity should be **O(log (m + n))**.

## Examples

### Example 1

**Input**:

```plaintext
nums1 = [1, 3], nums2 = [2]
```

**Output**:

```plaintext
2.00000
```

**Explanation**:

- When you merge the two arrays, you get `[1, 2, 3]`.
- The median of `[1, 2, 3]` is `2`.

### Example 2

**Input**:

```plaintext
nums1 = [1, 2], nums2 = [3, 4]
```

**Output**:

```plaintext
2.50000
```

**Explanation**:

- When you merge the two arrays, you get `[1, 2, 3, 4]`.
- The median of `[1, 2, 3, 4]` is `(2 + 3) / 2 = 2.5`.

## Constraints

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10^6 <= nums1[i], nums2[i] <= 10^6`

## Approaches

- [Approach 1](./solution_1/) - Merge-Based Median Finding
