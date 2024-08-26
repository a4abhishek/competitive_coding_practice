package main

import "fmt"

func expandAroundCenter(str string, oddPalindrome bool, i int) (int, int) {
	s := i - 1
	e := i + 1

	if !oddPalindrome {
		s = i
	}

	rs := []rune(str)
	for s >= 1 && e < len(rs)-1 && rs[s-1] == rs[e+1] {
		s--
		e++
	}

	return s, e
}

func longestPalindrome(str string) string {
	if len(str) == 0 {
		return ""
	}

	beginning := 0 // Beginning of the palindrome
	end := 0       // End of the palindrome

	rs := []rune(str)
	for i := 0; i < len(rs)-1; i++ {
		var s, e int

		if i > 0 && rs[i-1] == rs[i+1] {
			// odd palindrome
			// fmt.Printf("odd palindrome center=%d\n", i)
			s, e = expandAroundCenter(str, true, i)

			if e-s > end-beginning {
				beginning = s
				end = e
			}
		}

		if rs[i] == rs[i+1] {
			// even palindrome
			// fmt.Printf("even palindrome center=%d\n", i)
			s, e = expandAroundCenter(str, false, i)

			if e-s > end-beginning {
				beginning = s
				end = e
			}
		}
	}

	// Handling palindrome of length 1. Example: abc
	if end == 0 {
		return string(rs[0])
	}

	// fmt.Printf("%s ---> beginning: %d, end: %d, length: %d\n", str, beginning, end, end-beginning+1)
	return string(rs[beginning : end+1])
}

func main() {
	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("bab"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("ccc"))
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("abc"))
}
