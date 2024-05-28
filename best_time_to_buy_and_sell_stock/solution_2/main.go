package main

import "fmt"

func maxProfit(prices []int) int {
	maxAfterThis := make([]int, len(prices))

	currentMax := -1
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > currentMax {
			currentMax = prices[i]
		}

		maxAfterThis[i] = currentMax
	}

	maxProfit := -1
	for i, price := range prices {
		if maxAfterThis[i]-price > maxProfit {
			maxProfit = maxAfterThis[i] - price
		}
	}

	return maxProfit
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		prices := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&prices[j])
		}

		result := maxProfit(prices)
		fmt.Println(result)
	}
}
