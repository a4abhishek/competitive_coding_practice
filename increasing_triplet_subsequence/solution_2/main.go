package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt

	for _, n := range nums {
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}

	return false
}

func main() {
	// nums := []int{5, 1, 6, 2, 4}
	// nums := []int{4, 5, 6, 1, 2}
	// nums := []int{1, 1, 1}
	nums := []int{4, 5, 1, 6, 2}

	fmt.Println(increasingTriplet(nums))
}
