package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	maxSum := math.MinInt

	sum := 0
	s, e := 0, 0
	for e < len(nums) {
		sum += nums[e]

		if sum > maxSum {
			maxSum = sum
		}

		e++

		if e == len(nums) {
			s++
			e = s
			sum = 0
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, 1}))
}
