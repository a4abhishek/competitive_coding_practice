package main

import "fmt"

func expandAroundCenter(str string, start, end int) (int, int) {
	rs := []rune(str)
	for start >= 0 && end < len(rs) && rs[start] == rs[end] {
		start--
		end++
	}

	return start + 1, end - 1
}

func longestPalindrome(str string) string {
	if len(str) == 0 {
		// Without this check, the function might return the null character `"\x00"`,
		// which represents a non-existent value in the string.
		return ""
	}

	start := 0 // Beginning of the palindrome
	end := 0   // End of the palindrome

	rs := []rune(str)
	for i := 0; i < len(rs)-1; i++ {
		// odd palindrome
		s, e := expandAroundCenter(str, i, i)
		if e-s > end-start {
			start, end = s, e
		}

		// even palindrome
		s, e = expandAroundCenter(str, i, i+1)
		if e-s > end-start {
			start, end = s, e
		}
	}

	// fmt.Printf("%s ---> start: %d, end: %d, length: %d ---> %s\n", str, start, end, end-start+1, string(rs[start : end+1]))
	return string(rs[start : end+1])
}

func main() {
	fmt.Println(longestPalindrome(""))      // Expected: ""
	fmt.Println(longestPalindrome("a"))     // Expected: "a"
	fmt.Println(longestPalindrome("bab"))   // Expected: "bab"
	fmt.Println(longestPalindrome("bb"))    // Expected: "bb"
	fmt.Println(longestPalindrome("ccc"))   // Expected: "ccc"
	fmt.Println(longestPalindrome("aaaa"))  // Expected: "aaaa"
	fmt.Println(longestPalindrome("cbbd"))  // Expected: "bb"
	fmt.Println(longestPalindrome("babad")) // Expected: "bab" or "aba"
	fmt.Println(longestPalindrome("abc"))   // Expected: "a", "b", or "c"
}
