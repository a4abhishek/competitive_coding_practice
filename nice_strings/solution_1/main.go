package main

import "fmt"

const MOD = 1000000007

// number of possible nice strings of length n ending with each character.
// dp[i]['L'] be the number of strings of length i that end with 'L'
var dp = map[int]map[rune]int{}

// countNiceStringsEndsWith(n, r) calculates the number of nice strings of length n that end with character r ('L' or 'R').
func countNiceStringsEndsWith(n int, r rune) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		if r == 'L' {
			return 1
		}

		return 2
	}

	if c, ok := dp[n][r]; ok {
		return c
	}

	c := 0
	switch r {
	case 'L':
		c += countNiceStringsEndsWith(n-1, 'R')
	case 'R':
		c += countNiceStringsEndsWith(n-1, 'L')
		c %= MOD
		c += countNiceStringsEndsWith(n-1, 'R')
	}

	c %= MOD
	if _, ok := dp[n]; !ok {
		dp[n] = map[rune]int{}
	}
	dp[n][r] = c
	return c
}

func calculateSumOfSquares(n int) int {
	c := 0

	for i := 1; i <= n; i++ {
		cn := 0
		cn += countNiceStringsEndsWith(i, 'L')
		cn += countNiceStringsEndsWith(i, 'R')

		// fmt.Printf("length=%d, count=%d\n", i, cn)
		c += cn * cn
		c %= MOD
	}

	return c
}

func main() {
	// fmt.Println(calculateSumOfSquares(1))
	// fmt.Println(calculateSumOfSquares(2))
	// fmt.Println(calculateSumOfSquares(3))
	// fmt.Println(calculateSumOfSquares(4))
	fmt.Println(calculateSumOfSquares(5))
}
