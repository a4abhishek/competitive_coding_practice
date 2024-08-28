# Count Valid Subarrays

## Problem Statement

You are given an array of integers. A subarray is considered *valid* if the sum of its elements is not zero. Your task is to determine how many valid subarrays exist in the given array.

### Examples

1. **Example 1:**
   - **Input:** `[1, -1, 2, 3]`
   - **Output:** `9`
   - **Explanation:**
     - There are 10 subarrays in total: `[1]`, `[1, -1]`, `[1, -1, 2]`, `[1, -1, 2, 3]`, `[-1]`, `[-1, 2]`, `[-1, 2, 3]`, `[2]`, `[2, 3]`, `[3]`.
     - The only subarray with sum zero is `[1, -1]`.
     - Hence, there are 9 valid subarrays.

2. **Example 2:**
   - **Input:** `[0, 0, 0]`
   - **Output:** `0`
   - **Explanation:** All possible subarrays have a sum of zero, so the count of valid subarrays is `0`.

### Constraints

- The input array can have up to $10^5$ elements.
- Each element in the array is an integer within the range `[-10^9, 10^9]`.

## Approach (Prefix Sum Technique)

### Optimal Solution

To solve this problem optimally, we can use a variant of [Kadan's Algorithm](../maximum_subarray/README.md#kadanes-algorithm) in combination with a hash map to efficiently count the number of valid subarrays.

### Steps:

1. **Prefix Sum Calculation**:
   - We compute the prefix sum as we iterate through the array.
   - The prefix sum at index `i` is the sum of all elements from the start of the array up to index `i`.

2. **Using a Hash Map to Track Prefix Sums**:
   - We use a hash map to store the frequency of each prefix sum encountered during the iteration.
   - The key insight is that if the same prefix sum appears more than once, it means there is a subarray with a sum of zero between those indices. We want to exclude these from the count of valid subarrays.

3. **Count Valid Subarrays**:
   - As we iterate through the array, for each prefix sum, we check if it has appeared before in the hash map. If it has, the corresponding subarray(s) will have a sum of zero.
   - Subtract the count of subarrays with a sum of zero from the total number of subarrays to get the number of valid subarrays.

### Complexity Analysis

- **Time Complexity**: `O(n)`, where `n` is the number of elements in the array. We only make a single pass through the array, and all operations involving the hash map (insertion, lookup) take constant time.
- **Space Complexity**: `O(n)`, because we use a hash map to store the prefix sums and their frequencies.

### Example Walkthrough

Consider the array `[1, -1, 2, 3]`:

- **Step 1: Compute Prefix Sums**:
  - Prefix sums are `[1, 0, 2, 5]`.

- **Step 2: Track Prefix Sums Using Hash Map**:
  - Initially, the hash map contains `{0: 1}` (prefix sum 0 is seen before we start).
  - As we iterate:
    - At index 0: Prefix sum is `1`, hash map becomes `{0: 1, 1: 1}`.
    - At index 1: Prefix sum is `0`, hash map becomes `{0: 2, 1: 1}` (indicating a subarray with a sum of zero).
    - At index 2: Prefix sum is `2`, hash map becomes `{0: 2, 1: 1, 2: 1}`.
    - At index 3: Prefix sum is `5`, hash map becomes `{0: 2, 1: 1, 2: 1, 5: 1}`.

- **Step 3: Count Valid Subarrays**:
  - The only invalid subarray is `[1, -1]` (corresponding to the repeated prefix sum `0`).
  - Total subarrays: 10, invalid subarrays: 1, valid subarrays: `9`.

### Example Implementation in Python

Here’s an efficient Python implementation of the solution:

```python
def count_valid_subarrays(arr):
    prefix_sum = 0
    prefix_sum_map = {0: 1}  # Initialize with 0 to handle subarrays that sum to zero starting from the beginning
    invalid_subarray_count = 0
    
    for num in arr:
        prefix_sum += num
        
        if prefix_sum in prefix_sum_map:
            # This means there is a zero-sum subarray ending here
            invalid_subarray_count += prefix_sum_map[prefix_sum]
        else:
            prefix_sum_map[prefix_sum] = 0
        
        prefix_sum_map[prefix_sum] += 1
    
    total_subarrays = len(arr) * (len(arr) + 1) // 2
    valid_subarrays = total_subarrays - invalid_subarray_count
    
    return valid_subarrays

# Example usage
arr = [2, 1, -1, 5]
result = count_valid_subarrays(arr)
print(result)  # Output: 9
```

### Conclusion

This problem requires counting the number of valid subarrays in an array where the sum of the subarray is not zero. By utilizing a prefix sum and hash map approach, the solution is both time-efficient (`O(n)`) and space-efficient, making it suitable for large inputs.
