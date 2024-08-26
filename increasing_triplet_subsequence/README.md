# [LeetCode 334: Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence)

## Problem Statement

Given an integer array `nums`, return `true` if there exists a triple of indices `(i, j, k)` such that `i < j < k` and `nums[i] < nums[j] < nums[k]`. If no such indices exist, return `false`.

### Example 1:
- **Input:** nums = [1,2,3,4,5]
- **Output:** true
- **Explanation:** Any triplet where `i < j < k` is valid.

### Example 2:
- **Input:** nums = [5,4,3,2,1]
- **Output:** false
- **Explanation:** No triplet exists.

### Example 3:
- **Input:** nums = [2,1,5,0,4,6]
- **Output:** true
- **Explanation:** The triplet `(3, 4, 5)` is valid because `nums[3] == 0 < nums[4] == 4 < nums[5] == 6`.

### Constraints:
- `1 <= nums.length <= 5 * 10^5`
- `-2^31 <= nums[i] <= 2^31 - 1`

### Follow-up:
Could you implement a solution that runs in `O(n)` time complexity and `O(1)` space complexity?

## Approach

To solve this problem efficiently, we need to find a solution that can identify whether there exists an increasing triplet subsequence in the array in linear time and constant space.

### Solution

The key is to maintain two variables, `first` and `second`, to store the smallest and second smallest values found so far, respectively:

1. **Initialize `first` and `second`:**
   - Start by setting both `first` and `second` to a large value.

2. **Iterate through the array:**
   - For each element `num` in the array:
     - If `num` is smaller than or equal to `first`, update `first` with `num`.
     - Otherwise, if `num` is smaller than or equal to `second`, update `second` with `num`.
     - If `num` is greater than `second`, it means we have found a triplet where `first < second < num`, so we return `true`.

3. **Return `false` if no triplet is found:**
   - If the loop completes without finding a valid triplet, return `false`.

### Time Complexity

The solution runs in `O(n)` because we only traverse the array once.

### Space Complexity

The space complexity is `O(1)` because we only use a constant amount of extra space.

## Example Implementation in Python

```python
def increasingTriplet(nums):
    first = second = float('inf')
    
    for num in nums:
        if num <= first:
            first = num
        elif num <= second:
            second = num
        else:
            return True
    
    return False
```

### Clarifying a confusion

In this solution, the variable `second` carries an implicit guarantee: it indicates that there was a number before it that is smaller, even if `first` has since been updated to a different, smaller value that comes after `second`.

Let's break it down with an example:

- Consider the array `nums = [4, 5, 1, 6, 2]`.
  
When the loop runs:

1. `first` starts as 4, and then `second` becomes 5. This tells us that 4 (which is smaller than 5) exists before 5 in the sequence.
2. Next, `first` updates to 1 because 1 is smaller than 4. However, `second` remains 5. Even though `first` is now 1, which comes after 5, the fact that `second` is still 5 means we’ve already found a number (4) that was smaller and appeared before 5.

So, when we encounter the number 6, we can confidently say it is the third number in an increasing triplet (because we know 4 < 5 < 6, even though `first` was updated to 1 later on).

<!-- My first draft. Could be repetitive but still useful --
In this solution, second intrinsically contains the fact that there is a number before second which is smaller than it, even if first is now updated to some other value which comes after the current value of second.

For example:
nums: [4, 5, 1, 6, 2]

when the loop runs, first becomes 4, second becomes 5 (which intrinsically suggests that a number before 5 exists which is smaller than 5)

But after that first becomes 1, while second is still 5 which is before the second in this case. But just the fact that we had assigned 5 in second says that there is a number like 4 which comes before 5. So, when we encounter 6 we can safely say that it is the third number.
-->
