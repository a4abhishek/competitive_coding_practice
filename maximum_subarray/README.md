# [LeetCode 53: Maximum Subarray](https://leetcode.com/problems/maximum-subarray)

## Problem Statement

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

### Example

**Example 1:**
- **Input:** `nums = [-2,1,-3,4,-1,2,1,-5,4]`
- **Output:** `6`
- **Explanation:** The subarray `[4,-1,2,1]` has the largest sum of `6`.

**Example 2:**
- **Input:** `nums = [1]`
- **Output:** `1`
- **Explanation:** The subarray `[1]` has the largest sum of `1`.

**Example 3:**
- **Input:** `nums = [5,4,-1,7,8]`
- **Output:** `23`
- **Explanation:** The subarray `[5,4,-1,7,8]` has the largest sum of `23`.

### Constraints
- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

## Approach

### Kadane's Algorithm

The problem of finding the maximum sum of a contiguous subarray can be efficiently solved using **Kadane's Algorithm**. This algorithm is designed to find the maximum sum subarray in linear time, making it optimal for this problem.

### How Kadane's Algorithm Works

The algorithm works by iterating through the array and maintaining two variables:
1. **`current_max`**: The maximum sum of the subarray that ends at the current position.
2. **`global_max`**: The maximum sum found so far among all subarrays considered.

### Iterative Process

For each element in the array, the algorithm decides whether to:
1. **Extend the current subarray**: Add the current element to the `current_max` (which represents extending the previous subarray).
2. **Start a new subarray**: Start a new subarray with just the current element, if the current element alone is larger than `current_max + current element`.

After updating `current_max`, the algorithm checks if `current_max` is greater than `global_max`. If it is, `global_max` is updated.

### Example Walkthrough: `[-3, 2, 1, -1, 4, -2, 1]`

Let's walk through Kadane's Algorithm with the array `[-3, 2, 1, -1, 4, -2, 1]` to show how it covers all possible subarrays, including smaller subarrays like `[2, 1]`, `[1]`, `[4]`, and `[1]` at the end.

- **Initial State**:
  - `current_max = -3`
  - `global_max = -3`

- **Step 1** (`num = 2`):
  - `current_max = max(2, -3 + 2) = 2`
  - `global_max = max(-3, 2) = 2`
  - The algorithm decides to start a new subarray with `2` because `2` alone is greater than `-1` (the sum of `current_max` and `num`).

- **Step 2** (`num = 1`):
  - `current_max = max(1, 2 + 1) = 3`
  - `global_max = max(2, 3) = 3`
  - The algorithm extends the current subarray, which now includes `[2, 1]`.

- **Step 3** (`num = -1`):
  - `current_max = max(-1, 3 + (-1)) = 2`
  - `global_max = max(3, 2) = 3`
  - The subarray `[2, 1]` is still the maximum, but it's temporarily reduced by `-1`.

- **Step 4** (`num = 4`):
  - `current_max = max(4, 2 + 4) = 6`
  - `global_max = max(3, 6) = 6`
  - The subarray is extended to include `[2, 1, -1, 4]`, which is now the maximum subarray.

- **Step 5** (`num = -2`):
  - `current_max = max(-2, 6 + (-2)) = 4`
  - `global_max = max(6, 4) = 6`
  - The subarray `[2, 1, -1, 4]` remains the maximum, but `current_max` is reduced.

- **Step 6** (`num = 1`):
  - `current_max = max(1, 4 + 1) = 5`
  - `global_max = max(6, 5) = 6`
  - The algorithm considers `[1]` and `[4, -2, 1]`, but `[2, 1, -1, 4]` remains the best subarray.

### Coverage of All Subarrays

1. **Small Subarrays**: The algorithm implicitly considers all possible subarrays by updating `current_max` with either the current element alone or extending the previous subarray. For example, when processing `2` and `1`, the algorithm considers subarrays `[2]`, `[1]`, and `[2, 1]`.

2. **Single Element Subarrays**: Single element subarrays like `[4]` or `[1]` are considered when `current_max` is reset to the current element, which happens when the current element is greater than `current_max + current element`.

3. **Resetting `current_max`**: When the algorithm encounters a large negative number, such as `-3` at the beginning, it starts a new subarray with the next positive number (like `2`), effectively covering subarrays that start after the negative sequence.

### Final Result

After the iteration, `global_max` will contain the maximum sum of any contiguous subarray within the array. For the example array `[-3, 2, 1, -1, 4, -2, 1]`, the maximum sum is `6`, corresponding to the subarray `[2, 1, -1, 4]`.

### Complexity Analysis

- **Time Complexity**: `O(n)`, where `n` is the length of the array. The algorithm only requires a single pass through the array.
- **Space Complexity**: `O(1)`, since it uses a constant amount of additional space (only a few variables).

### Example Implementation in Python

```python
def maxSubArray(nums):
    current_max = global_max = nums[0]
    
    for num in nums[1:]:
        current_max = max(num, current_max + num)
        global_max = max(global_max, current_max)
    
    return global_max
```

### Conclusion

Kadane's Algorithm provides an optimal solution to the Maximum Subarray problem by efficiently considering all possible subarrays, including small and single-element subarrays, within a linear time complexity and constant space complexity. Its simplicity and efficiency make it a fundamental algorithm in the study of dynamic programming and array manipulation.
