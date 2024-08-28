# Count Nearly Lucky Numbers in an Array

## Problem Statement

You are given an array of integers. A number is considered *lucky* if it consists of only the digits `{3, 4, 7}`. A number is called a *nearly lucky number* if the count of lucky digits in the number is also a lucky number.

Your task is to determine how many numbers in the given array are nearly lucky numbers.

### Examples

1. **Example 1:**
   - **Input:** `[4407, 1234, 777, 347]`
   - **Output:** `3`
   - **Explanation:** 
     - `4407` has 3 lucky digits (`4`, `4`, `7`), and `3` is a lucky number.
     - `1234` has 1 lucky digit (`4`), and `1` is not a lucky number.
     - `777` has 3 lucky digits (`7`, `7`, `7`), and `3` is a lucky number.
     - `347` has 3 lucky digits (`3`, `4`, `7`), and `3` is a lucky number.
     - Therefore, the array contains 3 nearly lucky numbers.

2. **Example 2:**
   - **Input:** `[1, 2, 5]`
   - **Output:** `0`
   - **Explanation:** None of the numbers have any lucky digits, so the count of nearly lucky numbers is `0`.

### Constraints

- The input array can have up to $10^5$ elements.
- Each element in the array is a non-negative integer, and each number can have up to $10^5$ digits.

## Approach

To solve this problem, follow these steps:

1. **Identify Lucky Digits:**
   - Lucky digits are `3`, `4`, and `7`.

2. **Count Lucky Digits for Each Number:**
   - For each number in the array, traverse through each digit and count how many of them are lucky digits.
   - Use memoization to store the number of lucky digits for each number to avoid redundant calculations.

3. **Check if the Count is a Lucky Number:**
   - After counting the lucky digits for a number, check if this count is itself a lucky number (`3`, `4`, or `7`).

4. **Count Nearly Lucky Numbers:**
   - Increment a counter for each number in the array that is a nearly lucky number.

5. **Return the Result:**
   - Return the count of nearly lucky numbers.

### Example Walkthrough

Consider the array `[4407, 1234, 777, 347]`:

- **4407**:
  - Digits: `4`, `4`, `0`, `7`
  - Count of lucky digits: `3` (Lucky digits: `4`, `4`, `7`)
  - `3` is a lucky number, so `4407` is a nearly lucky number.
  
- **1234**:
  - Digits: `1`, `2`, `3`, `4`
  - Count of lucky digits: `2` (Lucky digits: `3`, `4`)
  - `2` is not a lucky number, so `1234` is not a nearly lucky number.
  
- **777**:
  - Digits: `7`, `7`, `7`
  - Count of lucky digits: `3`
  - `3` is a lucky number, so `777` is a nearly lucky number.
  
- **347**:
  - Digits: `3`, `4`, `7`
  - Count of lucky digits: `3`
  - `3` is a lucky number, so `347` is a nearly lucky number.

In this example, there are 3 nearly lucky numbers.

### Example Implementation in Go

Here’s a Go implementation of the solution:

```python
# Dictionary to memoize the number of lucky digits for each number
memo = {}

def is_lucky_digit(digit):
    """Check if a digit is a lucky digit (3, 4, or 7)."""
    return digit in {'3', '4', '7'}

def count_lucky_digits(n):
    """Count the number of lucky digits in a number, with memoization."""
    if n in memo:
        return memo[n]
    
    count = 0
    original_n = n
    
    while n > 0:
        digit = n % 10
        if is_lucky_digit(str(digit)):
            count += 1
        n //= 10
    
    memo[original_n] = count
    return count

def is_lucky_number(count):
    """Check if the count is a lucky number (3, 4, or 7)."""
    return count in {3, 4, 7}

def count_nearly_lucky_numbers(arr):
    """Count how many numbers in the array are nearly lucky numbers."""
    nearly_lucky_count = 0
    
    for number in arr:
        lucky_digit_count = count_lucky_digits(number)
        if is_lucky_number(lucky_digit_count):
            nearly_lucky_count += 1
    
    return nearly_lucky_count

# Example usage
arr = [4407, 1234, 777, 347]
result = count_nearly_lucky_numbers(arr)
print(result)  # Output: 3
```

### Complexity Analysis

- **Time Complexity**: `O(n * m)`, where `n` is the number of elements in the array and `m` is the number of digits in the largest number. We use memoization to reduce redundant calculations, making the process more efficient.
- **Space Complexity**: `O(1)` (ignoring the memoization storage), as we only use a few variables to keep track of counts.
