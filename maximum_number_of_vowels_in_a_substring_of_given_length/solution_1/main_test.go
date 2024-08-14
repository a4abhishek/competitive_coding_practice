package main

import "testing"

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		// Basic cases
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},

		// Edge cases
		{"aeiou", 5, 5},                // Whole string is vowels
		{"abcdefg", 2, 1},              // Only one vowel
		{"bbbbbbbbbbbbbbbbbbbb", 5, 0}, // No vowels at all
		{"aeiou", 1, 1},                // Single-character substring, all vowels

		// Sliding window check
		{"aeiobcdaiueo", 5, 5}, // Check for a combination of vowels and non-vowels

		// All same characters
		{"aaaaaa", 3, 3}, // All vowels, k less than length
		{"bbbbbb", 3, 0}, // No vowels, all same characters
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			actual := maxVowels(tt.s, tt.k)
			if actual != tt.expected {
				t.Errorf("maxVowels(%s, %d) = %d; expected %d", tt.s, tt.k, actual, tt.expected)
			}
		})
	}
}
