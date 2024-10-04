# Full Merge Median Finder

The problem requires finding the median of two sorted arrays in **O(log (m + n))** time, which indicates that we cannot simply merge the two arrays and sort them to find the median (which would take **O(m + n)** time).

## Key Intuition

To achieve **O(log (m + n))** time complexity, we can approach the problem by **partitioning the arrays** and using **binary search** on the smaller array to find the median.

### Plan

1. **Partitioning the Arrays**:
   - We want to divide `nums1` and `nums2` such that:
     - Left partition contains the first half of elements,
     - Right partition contains the second half of elements,
     - Both partitions are of equal size or differ by 1 element (when the total number of elements is odd).

2. **Conditions for Partitioning**:
   - Use binary search to adjust the partitioning of the smaller array.
   - The key condition is that all elements on the left side of the partition in `nums1` and `nums2` must be less than or equal to all elements on the right side of the partition.

3. **Median Calculation**:
   - If the total number of elements is **odd**, the median is the maximum of the left partition.
   - If the total number of elements is **even**, the median is the average of the maximum of the left partition and the minimum of the right partition.

## Example Walkthrough

Consider `nums1 = [1, 3]` and `nums2 = [2]`:

1. The total number of elements is `3` (odd), so we need to partition the arrays such that one side contains `2` elements, and the other side contains `1`.

2. Using binary search, we find that the correct partition is:
   - Left side: `[1, 2]`
   - Right side: `[3]`

3. Since the total number of elements is odd, the median is the maximum of the left side, which is `2`.

## Formula for Median

The formula for the median depends on the total number of elements:

For an odd number of elements:

```plaintext
median = max(left_part)
```

For an even number of elements:

```plaintext
median = (max(left_part) + min(right_part)) / 2
```

## Time Complexity

- **O(log(min(m, n)))**: We apply binary search on the smaller of the two arrays, reducing the search space by half in each iteration.

## Space Complexity

- **O(1)**: We use only a constant amount of extra space.

## Python Code

```python
def findMedianSortedArrays(nums1, nums2):
    # Ensure nums1 is the smaller array
    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1

    m, n = len(nums1), len(nums2)
    half_len = (m + n + 1) // 2
    
    imin, imax = 0, m

    while imin <= imax:
        i = (imin + imax) // 2
        j = half_len - i

        if i < m and nums2[j-1] > nums1[i]:
            # Increase i
            imin = i + 1
        elif i > 0 and nums1[i-1] > nums2[j]:
            # Decrease i
            imax = i - 1
        else:
            # i is perfect
            if i == 0: max_of_left = nums2[j-1]
            elif j == 0: max_of_left = nums1[i-1]
            else: max_of_left = max(nums1[i-1], nums2[j-1])

            if (m + n) % 2 == 1:
                return max_of_left

            if i == m: min_of_right = nums2[j]
            elif j == n: min_of_right = nums1[i]
            else: min_of_right = min(nums1[i], nums2[j])

            return (max_of_left + min_of_right) / 2.0
```

### Explanation of Code

1. **Ensure `nums1` is smaller**: This simplifies the problem as we can always perform binary search on the smaller array, reducing complexity.

2. **Binary Search**:
   - We perform binary search on the smaller array (`nums1`) to find the correct partition.
   - Adjust the partition based on comparisons between elements on the left and right sides of the partition.

3. **Median Calculation**:
   - If the total number of elements is odd, return the maximum of the left partition.
   - If even, return the average of the maximum of the left partition and the minimum of the right partition.
