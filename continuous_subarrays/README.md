# [LeetCode 2762: Continuous Subarrays](https://leetcode.com/problems/continuous-subarrays/)

## Problem Statement

You are given a **0-indexed** integer array `nums`. A subarray of `nums` is called **continuous** if:

Let `i`, `i + 1, ..., j` be the indices in the subarray. Then, for each pair of indices `i1 <= i`, `i2 <= j`, it holds that:

```plaintext
0 <= |nums[i1] - nums[i2]| <= 2
```

Your task is to return the **total number of continuous subarrays**.

A subarray is a contiguous non-empty sequence of elements within an array.

## Examples

### Example 1

**Input:**

```plaintext
nums = [5, 4, 2, 4]
```

**Output:**

```plaintext
8
```

**Explanation:**

- Continuous subarrays of size 1: `[5]`, `[4]`, `[2]`, `[4]`.
- Continuous subarrays of size 2: `[5, 4]`, `[4, 2]`, `[2, 4]`.
- Continuous subarray of size 3: `[4, 2, 4]`.
- There are no valid subarrays of size 4 that meet the condition.

Thus, the total continuous subarrays: `4 + 3 + 1 = 8`.

### Example 2

**Input:**

```plaintext
nums = [1, 2, 3]
```

**Output:**

```plaintext
6
```

**Explanation:**

- Continuous subarrays of size 1: `[1]`, `[2]`, `[3]`.
- Continuous subarrays of size 2: `[1, 2]`, `[2, 3]`.
- Continuous subarray of size 3: `[1, 2, 3]`.

Thus, the total continuous subarrays: `3 + 2 + 1 = 6`.

## Constraints

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`

## Approaches

- Approach 1: [Brute Force with Nested Loops](./solution_1/)
- Approach 2: [Sparse Table and Binary Search](./solution_2/)
- Approach 3: [Sliding Window Approach](./solution_3/) (Best)
