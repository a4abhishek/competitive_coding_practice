package main

import "fmt"

func maxProfit(prices []int) int {
    maxAfterThis := make([]int, len(prices))

    currentMax := -1
    currentMaxIndex := len(prices)-1
    for i := len(prices)-1; i>=0; i-- {
        if prices[i] > currentMax {
            currentMax = prices[i]
            currentMaxIndex = i
        }

        maxAfterThis[i] = currentMaxIndex
    }

    buy := 0
    sell := 0
    maxProfit := -1
    for i, price := range prices {
        if prices[maxAfterThis[i]] - price > maxProfit {
            buy = i
            sell = maxAfterThis[i]
            maxProfit = prices[maxAfterThis[i]] - price
        }
    }

    return prices[sell] - prices[buy]
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
