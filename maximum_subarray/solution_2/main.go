package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	// Kadane's algorithm
	currentMax, globalMax := nums[0], nums[0]
	for _, n := range nums[1:] {
		currentMax = max(n, currentMax+n)
		globalMax = max(currentMax, globalMax)
	}
	return globalMax
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, 1}))
}
