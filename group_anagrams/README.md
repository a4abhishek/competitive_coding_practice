# [LeetCode 49: Group Anagrams](https://leetcode.com/problems/group-anagrams/)

## Problem Statement

Given an array of strings `strs`, group the anagrams together. You can return the answer in **any order**.

An anagram is a word or phrase formed by rearranging the letters of another, typically using all the original letters exactly once.

## Examples

### Example 1

**Input**:

```plaintext
strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
```

**Output**:

```plaintext
[["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
```

**Explanation**:

- There is no string in `strs` that can be rearranged to form `"bat"`, so it forms a group by itself: `["bat"]`.
- The strings `"nat"` and `"tan"` are anagrams and can be rearranged to form each other, so they form another group: `["nat", "tan"]`.
- Similarly, `"ate"`, `"eat"`, and `"tea"` are anagrams and can be rearranged to form each other, so they form another group: `["ate", "eat", "tea"]`.

### Example 2

**Input**:

```plaintext
strs = [""]
```

**Output**:

```plaintext
[[""]]
```

**Explanation**:

- There is only one empty string in the input, so it forms a group by itself.

### Example 3

**Input**:

```plaintext
strs = ["a"]
```

**Output**:

```plaintext
[["a"]]
```

**Explanation**:

- There is only one string `"a"`, so it forms a group by itself.

## Constraints

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Approaches

- Approach 1: [Frequency-Based Anagram Grouping](./solution_1/)
- Approach 2: [Sorting-Based Anagram Grouping](./solution_2/)
