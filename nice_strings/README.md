# Nice String Count and Fibonacci Sequence Relation

## Problem Statement

A string is called a *nice string* if it is composed only of the characters 'R' and 'L' and does not contain consecutive 'L' characters. 

Given a positive integer `N`, define `G(I)` as the number of nice strings of length `I`. The task is to calculate:

$$\sum_{i=1}^{n} G(i)^2 \mod (10^9 + 7)$$

### Example
- **Input:** I = 2
- **Nice Strings:** LL, LR, RL, RR
- **Valid Nice Strings:** LR, RL, RR
- **Output:** G(2) = 3 (since 'LL' is not a nice string)

Similarly, `G(1) = 2` as there are no consecutive 'L' characters in any of the two possible strings.

## Constraints

- `1 <= N <= 10^5`
- `Modulo = 10^9 + 7`

## Approach
- [Count permutations based on certain rule](./solution_1/)
- [Custom fibbonacci based approach (Reccursive)](./solution_2/)
- [Custom fibbonacci based approach (Iterative)](./solution_3/)
