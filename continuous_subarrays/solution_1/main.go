package main

import (
	"fmt"
	"math"
)

func continuousSubarrays(nums []int) int64 {
	subArraysCount := int64(0)

	var minNum, maxNum int
	for i := range nums {
		maxNum = math.MinInt
		minNum = math.MaxInt

		for j := i; j < len(nums); j++ {
			if nums[j] > maxNum {
				maxNum = nums[j]
			}

			if nums[j] < minNum {
				minNum = nums[j]
			}

			if maxNum-minNum > 2 {
				break
			}

			subArraysCount++
		}
	}

	return subArraysCount
}

func main() {
	fmt.Println(continuousSubarrays([]int{5, 4, 2, 4})) // Output: 8
	fmt.Println(continuousSubarrays([]int{1, 2, 3}))    // Output: 6
}
