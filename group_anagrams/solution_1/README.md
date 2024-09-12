# Frequency-Based Anagram Grouping

## Approach

This approach groups anagrams by comparing the frequency of characters in each string. Instead of sorting the string (as in the sorting-based approach), we create a frequency map (or hash map) for each string to represent the count of each character. Anagrams will have identical frequency maps, which can be used to group them together.

### Key Ideas

1. **Character Frequency Map**:
   - For each string, a map is built where the keys are characters, and the values are the counts of how often each character appears in the string. This character frequency map serves as a unique identifier for that string.

2. **Grouping by Frequency Map**:
   - We check if a string's frequency map matches any existing frequency map in the list of grouped anagrams. If it matches, it means the string is an anagram of one of the existing groups, and we add it to that group. If no match is found, a new group is created.

3. **Efficient Matching**:
   - For each new string, its frequency map is compared against the frequency maps of existing anagram groups. If a match is found, the string is added to the group. Otherwise, a new group is created.

### Python Code (for better understanding)

Here’s a Python implementation of the approach:

```python
class Anagram:
    def __init__(self, id_map, char_count, members):
        self.id_map = id_map      # Frequency map of characters
        self.char_count = char_count  # Total count of characters
        self.members = members    # List of words that belong to this anagram group

def group_anagrams(strs):
    anagrams = []

    # Iterate over each word
    for word in strs:
        cur_id = {}
        curr_char_count = 0

        # Build character frequency map for the word
        for char in word:
            cur_id[char] = cur_id.get(char, 0) + 1
            curr_char_count += 1

        matched = False

        # Check if the current word's frequency map matches any existing anagram group
        for anagram in anagrams:
            if curr_char_count != anagram.char_count:
                continue  # Different length, can't be an anagram

            matched = True
            # Compare frequency maps
            for char, count in anagram.id_map.items():
                if cur_id.get(char, 0) != count:
                    matched = False
                    break

            # If it matches, add the word to the anagram group
            if matched:
                anagram.members.append(word)
                break

        # If no matching anagram group is found, create a new one
        if not matched:
            anagrams.append(Anagram(cur_id, curr_char_count, [word]))

    # Collect the result
    return [anagram.members for anagram in anagrams]

# Example usage:
print(group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))
# Output: [['eat', 'tea', 'ate'], ['tan', 'nat'], ['bat']]
```

### How This Approach Works

1. **Create Frequency Map for Each Word**:
   - For each string in the input array, build a character frequency map (a dictionary in Python or map in Go), where the keys are the characters in the string and the values are the count of each character.

2. **Compare with Existing Groups**:
   - We loop through the list of existing anagram groups and compare the frequency map of the current string with the frequency maps of each group. If a match is found, the string is added to the group. A match is determined by comparing both the character counts and the total number of characters in the string.

3. **Add to Group or Create a New One**:
   - If no match is found, create a new group for the current string and store its frequency map.

4. **Return the Groups**:
   - After processing all the strings, the `members` list of each anagram group is returned as the result.

### Time Complexity

- The time complexity is **O(n * k)**, where `n` is the number of words and `k` is the average length of the words. This is because we build a character frequency map for each word in **O(k)** time and then compare it to existing anagram groups.

### Space Complexity

- The space complexity is **O(n * k)** for storing the frequency maps and the grouped anagrams, where `n` is the number of words and `k` is the average length of the words.
