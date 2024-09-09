package main

import "fmt"

const MOD = 1000000007

func calculateSumOfSquares(n int) int {
	if n == 1 {
		return 4
	}

	if n == 2 {
		return 9
	}

	gp, gn := 2, 3
	c := gp*gp + gn*gn
	for i := 3; i <= n; i++ {
		gp, gn = gn, gp+gn

		// fmt.Printf("length=%d, count=%d\n", i, gn)
		c += gn * gn
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
