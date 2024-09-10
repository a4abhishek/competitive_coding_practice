# [LeetCode 150: Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)

## Problem Statement

You are given an array of strings `tokens` that represents an arithmetic expression in [Reverse Polish Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation).

Your task is to evaluate the expression and return an integer representing the value of the expression.

### Note

- The valid operators are `'+'`, `'-'`, `'*'`, and `'/'`.
- Each operand may be an integer or another expression.
- Division between two integers always **truncates toward zero**.
- There will be no division by zero.
- The input represents a valid arithmetic expression in reverse Polish notation.
- The answer and all intermediate calculations are guaranteed to fit in a **32-bit integer**.

## Examples

### Example 1

**Input:**

```plaintext
tokens = ["2", "1", "+", "3", "*"]
```

**Output:**

```plaintext
9
```

**Explanation:**

```plaintext
((2 + 1) * 3) = 9
```

### Example 2

**Input:**

```plaintext
tokens = ["4", "13", "5", "/", "+"]
```

**Output:**

```plaintext
6
```

**Explanation:**

```plaintext
(4 + (13 / 5)) = 6
```

### Example 3

**Input:**

```plaintext
tokens = ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
```

**Output:**

```plaintext
22
```

**Explanation:**

```plaintext
((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
```

## Constraints

- `1 <= tokens.length <= 10^4`
- `tokens[i]` is either an operator: `"+"`, `"-"`, `"*"`, or `"/"`, or an integer in the range `[-200, 200]`.

## Approach

### Plan

1. **Stack-based Evaluation**:
   - Since Reverse Polish Notation (RPN) is postfix notation, operators appear after their operands. This makes it well-suited for evaluation using a stack.
   - Iterate over the input `tokens`:
     - If the token is a number, push it onto the stack.
     - If the token is an operator, pop two operands from the stack, apply the operator, and push the result back onto the stack.
   - **Note:** When popping, always pop the second operand **`b`** first, followed by the first operand **`a`**, as this is the order in which they were pushed. This is especially important for asymmetric operators like `-` and `/`.
   - The final value on the stack will be the result of the expression.

2. **Handling Division**:
   - Since division needs to truncate toward zero, use integer division in Python (`//`) for positive results and manually handle negative division cases.

### Example Walkthrough

For the input:

```plaintext
tokens = ["2", "1", "+", "3", "*"]
```

1. Push `2` to the stack.
2. Push `1` to the stack.
3. Encounter `+`, pop **`b = 1`** and **`a = 2`**, compute `2 + 1 = 3`, and push `3` to the stack.
4. Push `3` to the stack.
5. Encounter `*`, pop **`b = 3`** and **`a = 3`**, compute `3 * 3 = 9`, and push `9` to the stack.
6. The stack now contains `9`, which is the final result.

### Example Implementation in Python

```python
def evalRPN(tokens):
    stack = []
    
    for token in tokens:
        if token not in "+-*/":
            stack.append(int(token))
        else:
            b = stack.pop()  # Pop second operand first
            a = stack.pop()  # Pop first operand next
            if token == '+':
                stack.append(a + b)
            elif token == '-':
                stack.append(a - b)
            elif token == '*':
                stack.append(a * b)
            elif token == '/':
                # Handle division truncating toward zero
                stack.append(int(a / b))
    
    return stack[0]
```

### Complexity Analysis

- **Time Complexity:** `O(n)` where `n` is the length of the `tokens` list. Each token is processed once, either pushed to or popped from the stack.
- **Space Complexity:** `O(n)` for the stack, which stores intermediate results during the computation.
