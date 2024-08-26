package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"bab", "bab"},
		{"bb", "bb"},
		{"ccc", "ccc"},
		{"aaaa", "aaaa"},
		{"cbbd", "bb"},
		{"babad", "bab"},
		{"abc", "a"},
		{"racecar", "racecar"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := longestPalindrome(tt.input)
			if result != tt.expected && len(result) != len(tt.expected) {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
