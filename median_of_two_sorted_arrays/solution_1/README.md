# Merge-Based Median Finding

This approach uses the concept of merging two sorted arrays similar to the **merge step** of merge sort. The key idea is to efficiently find the median by merging the arrays only until the median is found, instead of merging the entire arrays.

## Steps

1. **Calculate Median Index**:
   - First, calculate the total length of the combined arrays (`nums1` and `nums2`).
   - If the total length is odd, the median is the element at the middle index.
   - If the total length is even, the median is the average of the two middle elements.

2. **Merge Until Median**:
   - Use two pointers, `i` for `nums1` and `j` for `nums2`, to keep track of the current position in each array.
   - Compare the elements at `nums1[i]` and `nums2[j]`, and append the smaller element to a "virtual merged array."
   - Continue moving the pointers forward until the merged array reaches the median index.

3. **Handle Odd and Even Cases**:
   - If the combined length is odd, return the element at the median index.
   - If the combined length is even, return the average of the two middle elements.

## Python Code

```python
def findMedianSortedArrays(nums1, nums2):
    total_len = len(nums1) + len(nums2)
    median_index = total_len // 2

    i, j, k = 0, 0, 0
    prev_median, median = 0, 0
    
    # Loop until we reach the median index
    while k <= median_index:
        prev_median = median
        
        if j >= len(nums2) or (i < len(nums1) and nums1[i] <= nums2[j]):
            median = nums1[i]
            i += 1
        else:
            median = nums2[j]
            j += 1
        
        k += 1

    # If total length is odd, return the middle element
    if total_len % 2 == 1:
        return float(median)
    
    # If total length is even, return the average of two middle elements
    return (prev_median + median) / 2.0
```

## Explanation of Code

1. **Calculate Median Index**:
   - `total_len` gives the combined length of both arrays.
   - `median_index` represents the position of the median, which we need to reach by merging the arrays.

2. **Merge Until the Median**:
   - We use two pointers (`i` and `j`) for `nums1` and `nums2`, respectively, and a loop that continues until we reach the `median_index`.
   - During each iteration, the smaller of the elements at `nums1[i]` and `nums2[j]` is selected, and the pointer of that array is moved forward.

3. **Handle Odd and Even Cases**:
   - If the total length is odd, the median is the middle element.
   - If the total length is even, the median is the average of the two middle elements (`prev_median` and `median`).

## Time Complexity

- **O(m + n)**: We are iterating through both arrays, but only until we reach the median index.

## Space Complexity

- **O(1)**: This approach uses a constant amount of space for pointers and variables.

## Example Walkthrough

For example, consider `nums1 = [1, 3]` and `nums2 = [2]`:

1. The combined length is `3`, so the median is at index `1` (because `3 // 2 = 1`).
2. Start with `i = 0`, `j = 0`, and compare `nums1[i] = 1` and `nums2[j] = 2`. Since `nums1[i]` is smaller, move the pointer `i` and set `median = 1`.
3. Continue comparing `nums1[i] = 3` and `nums2[j] = 2`. This time, `nums2[j]` is smaller, so set `median = 2` and move the pointer `j`.
4. Now that we’ve reached the `median_index`, return `2` as the median.
