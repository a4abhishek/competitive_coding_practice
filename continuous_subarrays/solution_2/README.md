# Continuous Subarrays Using Sparse Table and Binary Search

## Intuition

The problem boils down to counting the number of subarrays where the **maximum difference between any two elements** is at most 2. In other words, for each subarray, we want to ensure that the difference between the minimum and maximum element does not exceed 2.

## Approach

We can efficiently solve the problem using **sparse tables** and **binary search**:

### Key Ideas

1. **Sparse Table for Range Queries**:
   - A sparse table is a data structure that allows us to quickly query the **minimum** and **maximum** values in any subarray. This is useful because we need to repeatedly check the difference between the minimum and maximum element in subarrays of varying lengths.
   - The sparse table enables us to retrieve the minimum or maximum value in a range in **O(1)** time, after an **O(N log N)** preprocessing step.

2. **Binary Search to Find the Rightmost Index**:
   - For each index `i`, we perform **binary search** to find the rightmost index `r` such that the subarray from `i` to `r` satisfies the condition `max - min <= 2`. This reduces the number of subarrays we need to check for validity.

3. **Adding Subarrays**:
   - For every valid index `r` found using binary search, the number of valid subarrays starting at index `i` is `(r - i + 1)`. This is because every element from `i` to `r` can form a valid subarray.

---

### Step-by-Step Explanation with Example

Consider the input array:

```plaintext
nums = [5, 4, 2, 6]
```

We will show how the **sparse tables** are built and how **binary search** is used to find the rightmost valid subarray for each starting index.

#### Step 1: Build the Sparse Table

We create two sparse tables: one for the **minimum values** (`dpmn`) and one for the **maximum values** (`dpmx`). These tables will store the min and max values for subarrays of size `2^j`, where `j` ranges from `0` to `log(n)`.

- `dpmn[i][j]`: Stores the minimum value of the subarray starting at `i` of length `2^j`.
- `dpmx[i][j]`: Stores the maximum value of the subarray starting at `i` of length `2^j`.

**Sparse Table Construction:**

For `nums = [5, 4, 2, 6]`, the initial step (where `j = 0`) will just store the values of `nums` themselves, since a subarray of size `2^0 = 1` is just the element itself.

| Index `i` | `dpmn[i][0]`  | `dpmn[i][1]` | `dpmn[i][2]` | `dpmx[i][0]`  | `dpmx[i][1]` | `dpmx[i][2]` |
|-----------|---------------|--------------|--------------|---------------|--------------|--------------|
| 0         | 5             | 4            | 2            | 5             | 5            | 6            |
| 1         | 4             | 2            | N/A          | 4             | 6            | N/A          |
| 2         | 2             | 2            | N/A          | 2             | 6            | N/A          |
| 3         | 6             | N/A          | N/A          | 6             | N/A          | N/A          |

- For `j = 1` (subarrays of size `2`):
  - The min for each pair of consecutive elements is computed. For example, `dpmn[0][1] = min(dpmn[0][0], dpmn[1][0]) = min(5, 4) = 4`.
  - Similarly, the max is computed. For example, `dpmx[0][1] = max(dpmx[0][0], dpmx[1][0]) = max(5, 4) = 5`.

- For `j = 2` (subarrays of size `4`):
  - The min and max for the entire array are computed. For example, `dpmn[0][2] = min(dpmn[0][1], dpmn[2][1]) = min(4, 2) = 2`.

Now that the sparse table is built, we can efficiently query any subarray’s min or max in **O(1)** time.

---

#### Step 2: Perform Binary Search

For each index `i`, we will now use binary search to find the rightmost index `r` such that the subarray from `i` to `r` satisfies `max - min <= 2`.

1. **Start with `i = 0`:**

   - We want to find the rightmost `r` such that `max(nums[0:r]) - min(nums[0:r]) <= 2`.
   - Binary search between `lo = 0` and `hi = 3`.

   - **First Binary Search Step** (`mid = 2`):
     - Query the range `[0, 2]` using the sparse table:
       - `getmax(0, 2) = 5`
       - `getmin(0, 2) = 2`
       - Since `5 - 2 = 3 > 2`, we reduce the search space (`hi = mid - 1`).

   - **Second Binary Search Step** (`mid = 1`):
     - Query the range `[0, 1]` using the sparse table:
       - `getmax(0, 1) = 5`
       - `getmin(0, 1) = 4`
       - Since `5 - 4 = 1 <= 2`, we expand the search space (`lo = mid`).

   - The final valid index for `i = 0` is `r = 1`, so we add `(r - i + 1) = (1 - 0 + 1) = 2` subarrays.

2. **Next, for `i = 1`:**

   - Binary search between `lo = 1` and `hi = 3`.

   - **First Binary Search Step** (`mid = 2`):
     - Query the range `[1, 2]` using the sparse table:
       - `getmax(1, 2) = 4`
       - `getmin(1, 2) = 2`
       - Since `4 - 2 = 2 <= 2`, we expand the search space (`lo = mid`).

   - **Second Binary Search Step** (`mid = 3`):
     - Query the range `[1, 3]` using the sparse table:
       - `getmax(1, 3) = 6`
       - `getmin(1, 3) = 2`
       - Since `6 - 2 = 4 > 2`, we reduce the search space (`hi = mid - 1`).

   - The final valid index for `i = 1` is `r = 2`, so we add `(r - i + 1) = (2 - 1 + 1) = 2` subarrays.

3. **For `i = 2`:**

   - Binary search between `lo = 2` and `hi = 3`.

   - **First Binary Search Step** (`mid = 3`):
     - Query the range `[2, 3]` using the sparse table:
       - `getmax(2, 3) = 6`
       - `getmin(2, 3) = 2`
       - Since `6 - 2 = 4 > 2`, we reduce the search space (`hi = mid - 1`).

   - The final valid index for `i = 2` is `r = 2`, so we add `(r - i + 1) = (2 - 2 + 1) = 1` subarray.

4. **For `i = 3`:**
   - Only one element is considered (`nums[3] = 6`), and since a single element trivially satisfies the condition, we add `1` subarray.

---

### Final Count of Subarrays

After processing all indices, the total number of valid subarrays is `2 + 2 + 1 + 1 = 6`.

---

### Code Implementation (Python)

```python
import math

def continuousSubarrays(nums):
    n = len(nums)
    
    # Calculate maximum power of 2 less than or equal to n
    log = int(math.log2(n)) + 1
    
    # Initialize sparse tables for min and max values
    dpmn = [[float('inf')] * log for _ in range(n)]  # Sparse table for minimum values
    dpmx = [[float('-inf')] * log for _ in range(n)] # Sparse table for maximum values
    
    # Populate the initial values in the sparse tables
    for i in range(n):
        dpmn[i][

0] = nums[i]
        dpmx[i][0] = nums[i]
    
    # Build sparse tables
    j = 1
    while (1 << j) <= n:
        for i in range(n - (1 << j) + 1):
            dpmn[i][j] = min(dpmn[i][j - 1], dpmn[i + (1 << (j - 1))][j - 1])
            dpmx[i][j] = max(dpmx[i][j - 1], dpmx[i + (1 << (j - 1))][j - 1])
        j += 1
    
    # Helper functions to get the minimum and maximum values in a subarray
    def getmin(l, r):
        j = int(math.log2(r - l + 1))
        return min(dpmn[l][j], dpmn[r - (1 << j) + 1][j])
    
    def getmax(l, r):
        j = int(math.log2(r - l + 1))
        return max(dpmx[l][j], dpmx[r - (1 << j) + 1][j])
    
    ans = 0
    
    # Find the right index 'r' for each index 'i'
    for i in range(n):
        lo, hi = i, n - 1  # Binary search bounds
        
        # Perform binary search to find the rightmost index 'r'
        while lo < hi:
            mid = (lo + hi + 1) // 2
            if getmax(i, mid) - getmin(i, mid) <= 2:
                lo = mid
            else:
                hi = mid - 1
        
        # Add the count of valid subarrays starting from index 'i'
        ans += lo - i + 1
    
    return ans
```

---

### Time and Space Complexity

- **Time Complexity**:
  - **O(N log N)** for building the sparse table.
  - **O(N log N)** for the binary search for each index `i`.
  - Overall complexity: **O(N log N)**.
  
- **Space Complexity**: **O(N log N)** due to the sparse table.

---

### Problem Types Where This Approach Works

This approach can be generalized for problems where:

- You need to **query range min/max** values multiple times.
- Efficient handling of **range queries** is essential.
- **Binary search** combined with preprocessing (like sparse tables) can optimize the search for a valid range.
