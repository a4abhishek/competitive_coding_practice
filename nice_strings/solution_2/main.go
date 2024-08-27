package main

import "fmt"

const MOD = 1000000007

// number of possible nice strings of length n.
var dp = map[int]int{}

func countNiceStrings(n int) int {
	if n == 1 {
		return 2 // G(1) = 2
	}

	if n == 2 {
		return 3 // G(2) = 3
	}

	if c, ok := dp[n]; ok {
		return c % MOD
	}

	c := 0
	c += countNiceStrings(n - 1) // Append 'R' to any nice string of length n-1
	c %= MOD
	c += countNiceStrings(n - 2) // Append 'RL' to a nice string of length n-2
	c %= MOD

	dp[n] = c
	return c % MOD
}

func calculateSumOfSquares(n int) int {
	c := 0

	for i := 1; i <= n; i++ {
		cn := countNiceStrings(i)

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
