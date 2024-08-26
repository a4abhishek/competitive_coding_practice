package main

import "fmt"

// increasingSlice returns true if there are increasing steps of length "depth" starting 0th element. "depth" includes 0th element.
func increasingSlice(nums []int, depth int) bool {
	if depth < 0 || len(nums) < depth {
		return false
	}

	if len(nums) > 0 && depth == 1 {
		return true
	}

	for i := 1; i < len(nums); i++ {
		if nums[0] < nums[i] && increasingSlice(nums[i:], depth-1) {
			return true
		}
	}

	return false
}

func increasingTriplet(nums []int) bool {
	for i := range len(nums) {
		if increasingSlice(nums[i:], 3) {
			return true
		}
	}

	return false
}

func main() {
	// nums := []int{5, 1, 6, 2, 4}
	nums := []int{4, 5, 6, 1, 2}

	fmt.Println(increasingTriplet(nums))
}
