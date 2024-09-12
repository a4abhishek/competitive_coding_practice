# Continuous Subarrays Using Brute Force with Nested Loops

## Intuition

The idea behind this approach is straightforward: for each starting index `i` in the array, we try to extend the subarray as far as possible while ensuring that the difference between the maximum and minimum elements in the subarray is no more than 2.

For each valid subarray, we increment a count of valid subarrays. If at any point the difference between the maximum and minimum elements in the subarray exceeds 2, we stop extending the subarray.

## Approach

1. **Initialize Counters**:
   - Use two variables `minNum` and `maxNum` to track the minimum and maximum elements in the current subarray.
   - For every starting index `i` in the array, we reset `minNum` to infinity and `maxNum` to negative infinity.

2. **Iterate Over the Array**:
   - For each starting index `i`, try to extend the subarray by including more elements to the right (index `j`).
   - Update `minNum` and `maxNum` as you include more elements into the subarray.
   - If at any point the condition `maxNum - minNum > 2` is violated, break out of the inner loop.

3. **Count Valid Subarrays**:
   - Every time the condition is satisfied, increment the count of valid subarrays by 1.
   - Continue this process until all possible subarrays have been checked.

### Python Implementation

Here is a simplified Python version of the code for better understanding:

```python
def continuous_subarrays(nums):
    subarrays_count = 0

    for i in range(len(nums)):
        max_num = float('-inf')  # Initialize max to a very small number
        min_num = float('inf')   # Initialize min to a very large number

        for j in range(i, len(nums)):
            max_num = max(max_num, nums[j])  # Update max element in the subarray
            min_num = min(min_num, nums[j])  # Update min element in the subarray

            # If the difference exceeds 2, stop extending this subarray
            if max_num - min_num > 2:
                break

            # Count this valid subarray
            subarrays_count += 1

    return subarrays_count

# Test the function
print(continuous_subarrays([5, 4, 2, 4]))  # Output: 8
print(continuous_subarrays([1, 2, 3]))     # Output: 6
```

### Step-by-Step Walkthrough (Example)

Consider the input `nums = [5, 4, 2, 4]`:

1. **For `i = 0` (starting from element `5`)**:
   - Subarray `[5]`: valid, count = 1
   - Subarray `[5, 4]`: valid, count = 2
   - Subarray `[5, 4, 2]`: invalid, because `5 - 2 = 3` exceeds 2. Stop.

2. **For `i = 1` (starting from element `4`)**:
   - Subarray `[4]`: valid, count = 3
   - Subarray `[4, 2]`: valid, count = 4
   - Subarray `[4, 2, 4]`: valid, count = 5

3. **For `i = 2` (starting from element `2`)**:
   - Subarray `[2]`: valid, count = 6
   - Subarray `[2, 4]`: valid, count = 7

4. **For `i = 3` (starting from element `4`)**:
   - Subarray `[4]`: valid, count = 8

Total valid subarrays: `8`.

### Time Complexity

- The time complexity of this approach is **O(n²)**, where `n` is the length of the array. This is because, for each element `i`, we check all possible subarrays starting at `i` by iterating over the remaining elements in the array.

### Space Complexity

- The space complexity is **O(1)**, as we only use a few variables (`minNum`, `maxNum`, `subarrays_count`) to store intermediate values, and no additional data structures are required.
