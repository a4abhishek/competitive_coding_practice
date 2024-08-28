package main

import "fmt"

// Map of lucky digits
var luckyDigits = map[int]bool{
	3: true,
	4: true,
	7: true,
}

// isLuckyDigit checks if a digit is lucky
func isLuckyDigit(digit int) bool {
	return luckyDigits[digit]
}

// dp is a memoization map to store the number of lucky digits for a given number
var dp = map[int]int{} // dp[347] will return the number of lucky digits in the number 347 which is 3

// luckyDigitCount counts the number of lucky digits in a number
func luckyDigitCount(n int) int {
	if n < 10 {
		if isLuckyDigit(n) {
			return 1
		}

		return 0
	}

	if c, ok := dp[n]; ok {
		return c
	}

	c := 0
	for n > 0 {
		if isLuckyDigit(n % 10) {
			c++
		}

		n /= 10
		c += luckyDigitCount(n)
	}

	dp[n] = c
	return c
}

// countNearlyLuckyNumbers counts the number of nearly lucky numbers in an array
func countNearlyLuckyNumbers(arr []int) int {
	nearlyLuckyCount := 0

	for _, n := range arr {
		if isLuckyDigit(luckyDigitCount(n)) {
			nearlyLuckyCount++
		}
	}

	return nearlyLuckyCount
}

func main() {
	fmt.Println(countNearlyLuckyNumbers([]int{4407, 1234, 777, 347}))
}
