# Continuous Subarrays Using Sliding Window Approach

The problem can be efficiently solved using a **sliding window** technique. This approach works by expanding and shrinking a window of elements (subarray) to ensure that the condition `0 <= |nums[i1] - nums[i2]| <= 2` is satisfied for all elements within the window.

Here's how it works step-by-step:

1. **Expand the window** by adding elements from the right until the window becomes invalid (i.e., there exists a pair of elements where the difference is greater than `2`).

2. Once the window becomes invalid, **shrink the window** from the left until the condition is restored.

3. For every valid window (ending at `right`), the number of valid subarrays that can be formed is equal to the length of the window: `(right - left + 1)`. This is because every element in the window can serve as the starting point of a subarray, and the subarray can extend up to `right`.

## Code Implementation (Python)

```python
from collections import deque

def count_continuous_subarrays(nums):
    left = 0
    total_subarrays = 0
    max_deque = deque()
    min_deque = deque()

    for right in range(len(nums)):
        # Maintain the max deque (store indices in decreasing order of values)
        while max_deque and nums[max_deque[-1]] <= nums[right]:
            max_deque.pop()
        max_deque.append(right)
        
        # Maintain the min deque (store indices in increasing order of values)
        while min_deque and nums[min_deque[-1]] >= nums[right]:
            min_deque.pop()
        min_deque.append(right)
        
        # Shrink the window from the left to ensure |nums[i1] - nums[i2]| <= 2
        while nums[max_deque[0]] - nums[min_deque[0]] > 2:
            left += 1
            if max_deque[0] < left:
                max_deque.popleft()
            if min_deque[0] < left:
                min_deque.popleft()

        # Calculate the number of valid subarrays ending at 'right'
        total_subarrays += (right - left + 1)

    return total_subarrays
```

## Intuitive Explanation of Subarray Counting

For each window that ends at index `right`, the number of valid subarrays that can be formed is straightforward. If the window length is `n` (i.e., from `left` to `right`), you can form `n` subarrays:

- Start from `left`, and end at `right`
- Start from `left + 1`, and end at `right`
- Continue until you start at `right`, forming a subarray of size 1.

So for every valid window, the total subarrays are computed as `(right - left + 1)`.

## Maintaining the Maximum and Minimum Values

To efficiently track the largest and smallest values in the sliding window, we use two **deques**:

- **max_deque**: This deque stores indices of elements in decreasing order of their values, so that the front always holds the maximum value of the current window.
- **min_deque**: This deque stores indices in increasing order of their values, ensuring that the front always holds the minimum value of the current window.

By maintaining these deques, we can check if the difference between the maximum and minimum value exceeds `2` in constant time.

## Deque Properties (reminder)

The deque used in this algorithm has the following important properties:

- **Index order is always incremental**: As we slide the window from left to right, the indices in the deque naturally progress from lower to higher.
- **Values follow the nature of the deque**: For `max_deque`, indices are added in decreasing order of their corresponding values (i.e., the highest value comes first). Similarly, for `min_deque`, indices are added in increasing order of their values (i.e., the smallest value comes first).
- **Deque only stores values valid within the sliding window**: If an index falls outside the current sliding window, it is removed from the front of the deque.

## Why We Remove Elements from the Deque’s Rear

We remove elements from the rear of the deque when a new element comes in that is greater (for `max_deque`) or smaller (for `min_deque`). Here’s why:

- Consider the array `nums = [5, 4, 2, 3]`, and let’s assume the current sliding window contains the indices `[0, 1, 2]`, corresponding to the elements `[5, 4, 2]`.
- Now, when we try to add the element `3` at index `3`, it is greater than the element `2` at index `2`.
  - We can remove index `2` from the deque because `2` will never be useful again. If the window shrinks to remove index `1`, index `3` (value `3`) will always be higher, making index `2` unnecessary. Even if the window shrinks further, index `3` will still be part of the window, and index `2` will remain irrelevant.
  - This ensures that the deque always holds only the most relevant elements for determining the minimum or maximum value in the current window.

Thus, the deques store **indices in increasing order** for `min_deque` and **decreasing order** for `max_deque`, and always retain only those values that are still valid in the current window.

## Example Walkthrough

Let’s walk through an example with `nums = [5, 4, 2, 4]`:

1. Start with an empty window.
2. Expand the window by adding elements to the right while keeping the condition `0 <= |nums[i1] - nums[i2]| <= 2` satisfied.
3. If the condition is violated (i.e., the difference between the largest and smallest elements exceeds 2), move the left boundary (`left`) until the window becomes valid again.
4. For each valid window ending at `right`, count the number of subarrays formed by this window using `(right - left + 1)`.

## Problem Types Where This Approach Works

This sliding window approach using deques can be applied to other problems where:

- You need to maintain a range or condition within a sliding window.
- You want to efficiently track the minimum or maximum element in a subarray/window.
- For example, problems related to:
  - **Sliding window maximum/minimum**.
  - **Find the longest subarray with a condition on the difference between elements**.
  - **Maintaining specific ranges of values or sums** in a subarray.

## Complexity Analysis

- **Time Complexity**: `O(n)` where `n` is the length of the array. Each element is added and removed from the deque at most once, making the overall time complexity linear.
  
- **Space Complexity**: `O(n)` due to the space used by the deques.
