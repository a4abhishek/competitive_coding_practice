# [LeetCode 5: Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring)

## Problem Statement

Given a string `s`, return the longest palindromic substring in `s`.

### Example 1

- **Input:** s = "babad"
- **Output:** "bab"
- **Explanation:** "aba" is also a valid answer.

### Example 2

- **Input:** s = "cbbd"
- **Output:** "bb"

### Constraints

- `1 <= s.length <= 1000`
- `s` consists of only digits and English letters.

## Approach

To solve this problem, we need to identify the longest substring in `s` that reads the same backward as forward. There are several approaches to solve this problem, including brute force, dynamic programming, and expanding around the center.

### Solution 1: Expand Around Center (Optimal Solution)

This approach considers every character (and the gap between every two characters) as the potential center of a palindrome and expands outwards as long as the characters on both sides are equal.

#### Steps

1. **Initialize variables**
   - `start` and `end` to keep track of the starting and ending indices of the longest palindromic substring found.

2. **Expand around each center**:
   - For each character and the gap between characters in the string, consider it as the center of a potential palindrome.
   - Expand outward from the center as long as the characters at both ends match.
   - Update `start` and `end` if a longer palindrome is found.

3. **Return the longest palindrome**:
   - After checking all possible centers, return the substring starting at `start` and ending at `end`.

### Time Complexity

The time complexity of this approach is `O(n^2)`, where `n` is the length of the string. This is because we are expanding around each character (and the gap between each character) in the string.

### Space Complexity

The space complexity is `O(1)` since we are using a constant amount of extra space.

## Example Implementation in Python

```python
def expandAroundCenter(s, left, right):
    while left >= 0 and right < len(s) and s[left] == s[right]:
        left -= 1
        right += 1
    return left + 1, right - 1


def longestPalindrome(s):
    start, end = 0, 0
    for i in range(len(s)):
        l1, r1 = expandAroundCenter(s, i, i)
        l2, r2 = expandAroundCenter(s, i, i + 1)
        
        if r1 - l1 > end - start:
            start, end = l1, r1
        if r2 - l2 > end - start:
            start, end = l2, r2
    
    return s[start:end + 1]
```

### Other Approaches

1. **Dynamic Programming**:
   - This approach involves building a DP table where `dp[i][j]` is `True` if the substring `s[i:j+1]` is a palindrome.
   - The time and space complexity for this approach is `O(n^2)`.

2. **Brute Force**:
   - Check all possible substrings and verify if they are palindromes.
   - This approach has a time complexity of `O(n^3)` and is not efficient for large strings.
