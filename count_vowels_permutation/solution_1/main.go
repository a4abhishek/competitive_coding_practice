package main

import "fmt"

const MOD = 1000000007

// number of possible valid strings of length n ending with each vowel
// dp[i]['a'] be the number of strings of length i that end with the vowel 'a'
var dp = map[int]map[rune]int{}

func countPermutationEndswith(n int, r rune) int {
	if n == 1 {
		return 1
	}

	if c, ok := dp[n][r]; ok {
		return c
	}

	c := 0
	switch r {
	case 'a':
		c += countPermutationEndswith(n-1, 'e')
		c %= MOD
		c += countPermutationEndswith(n-1, 'i')
		c %= MOD
		c += countPermutationEndswith(n-1, 'u')
	case 'e':
		c += countPermutationEndswith(n-1, 'a')
		c %= MOD
		c += countPermutationEndswith(n-1, 'i')
	case 'i':
		c += countPermutationEndswith(n-1, 'e')
		c %= MOD
		c += countPermutationEndswith(n-1, 'o')
	case 'o':
		c += countPermutationEndswith(n-1, 'i')
	case 'u':
		c += countPermutationEndswith(n-1, 'i')
		c %= MOD
		c += countPermutationEndswith(n-1, 'o')
	}

	if _, ok := dp[n]; !ok {
		dp[n] = map[rune]int{}
	}

	c %= MOD
	dp[n][r] = c
	return c
}

func countVowelPermutation(n int) int {
	if n == 1 {
		return 5
	}
	
	c := 0
	c += countPermutationEndswith(n, 'a')
	c %= MOD
	c += countPermutationEndswith(n, 'e')
	c %= MOD
	c += countPermutationEndswith(n, 'i')
	c %= MOD
	c += countPermutationEndswith(n, 'o')
	c %= MOD
	c += countPermutationEndswith(n, 'u')
	c %= MOD

	return c
}

func main() {
	n := 5

	fmt.Println(countVowelPermutation(n))
}
