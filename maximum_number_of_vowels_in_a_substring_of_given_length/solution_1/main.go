package main

import "fmt"

var isVowel = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func maxVowels(s string, k int) int {
	maxVowelCount := 0

	// Vowels in first k characters
	for i := range k {
		if isVowel[s[i]] {
			maxVowelCount++
		}
	}

	// Sliding window
	i := 0
	j := k - 1
	vowelCount := maxVowelCount
	for j < len(s) {
		// Slide left boundary
		if isVowel[s[i]] {
			vowelCount--
		}
		i++

		// Slide right boundary
		j++
		if j >= len(s) {
			break
		}

		if isVowel[s[j]] {
			vowelCount++
		}

		// Keep the max count
		if vowelCount > maxVowelCount {
			maxVowelCount = vowelCount
		}
	}

	// Return the maximum
	return maxVowelCount
}

func main() {
	input := "abciiidef"

	fmt.Println(maxVowels(input, 3))
}
