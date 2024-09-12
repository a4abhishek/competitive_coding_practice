# Sorting-Based Anagram Grouping

## Intuition

The key idea is that two words are **anagrams** if and only if they contain the same characters in the same frequency, regardless of their order. Thus, sorting the letters in a word will give the same result for all anagrams. This insight can be used to group words that are anagrams.

## Plan

1. **Sort Each Word**:
   - For each word in the input list, we can sort the characters in alphabetical order. Anagrams will result in the same sorted string.

2. **Use a Hash Map**:
   - Use a dictionary (hash map) where the key is the sorted version of the word, and the value is a list of words (anagrams) that match that sorted key.

3. **Group Anagrams**:
   - For each word, compute its sorted version and insert it into the dictionary under the corresponding sorted key.

4. **Return the Groups**:
   - The values of the dictionary represent the groups of anagrams. Return all the values as the result.

### Step-by-Step Walkthrough

1. **Input**: `["eat", "tea", "tan", "ate", "nat", "bat"]`
   - Start by initializing an empty dictionary.

2. **Processing the Strings**:
   - For `"eat"`, the sorted version is `"aet"`. Store `["eat"]` under the key `"aet"`.
   - For `"tea"`, the sorted version is `"aet"`. Append `"tea"` to the list under the key `"aet"`, resulting in `["eat", "tea"]`.
   - For `"tan"`, the sorted version is `"ant"`. Store `["tan"]` under the key `"ant"`.
   - For `"ate"`, the sorted version is `"aet"`. Append `"ate"` to the list under the key `"aet"`, resulting in `["eat", "tea", "ate"]`.
   - For `"nat"`, the sorted version is `"ant"`. Append `"nat"` to the list under the key `"ant"`, resulting in `["tan", "nat"]`.
   - For `"bat"`, the sorted version is `"abt"`. Store `["bat"]` under the key `"abt"`.

3. **Final Grouping**:
   - The dictionary now contains:

     ```plaintext
     {
       "aet": ["eat", "tea", "ate"],
       "ant": ["tan", "nat"],
       "abt": ["bat"]
     }
     ```

   - The grouped anagrams are the values of this dictionary:

     ```plaintext
     [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
     ```

## Python Code

```python
from collections import defaultdict

def group_anagrams(strs):
    anagram_map = defaultdict(list)  # Dictionary to store sorted word as key and list of anagrams as value

    for word in strs:
        # Sort the word and use it as the key
        sorted_word = ''.join(sorted(word))
        anagram_map[sorted_word].append(word)

    # Return the values (groups of anagrams) from the dictionary
    return list(anagram_map.values())

# Example usage:
print(group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))
# Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
```

## Time Complexity

- Sorting each word takes **O(k log k)** where `k` is the length of the word.
- Since there are `n` words, the total time complexity is **O(n * k log k)** where `n` is the number of words, and `k` is the average length of the words.

## Space Complexity

- The space complexity is **O(n * k)** where `n` is the number of words and `k` is the average length of each word. This is because we are storing the words in a hash map and also storing the sorted versions of the words as keys.
