package main

import "fmt"

func countValidSubarrays(arr []int) int {
	prefixSum := 0
	prefixSumMap := map[int]int{0: 1} // DP: Initialize with 0 to handle subarrays that sum to zero starting from the beginning
	invalidSubarrayCount := 0

	for _, num := range arr {
		prefixSum += num

		// Check if this prefix sum has been seen before
		if count, exists := prefixSumMap[prefixSum]; exists {
			invalidSubarrayCount += count
		}

		// Increment the count for this prefix sum in the map
		prefixSumMap[prefixSum] += 1
	}

	// Calculate total possible subarrays
	n := len(arr)
	totalSubarrays := n * (n + 1) / 2

	// Valid subarrays are total subarrays minus the invalid subarrays
	validSubarrays := totalSubarrays - invalidSubarrayCount

	return validSubarrays
}

func main() {
	fmt.Println(countValidSubarrays([]int{2, 1, -1, 5}))    // Output: 9
	fmt.Println(countValidSubarrays([]int{-2, 3, 3, 2, 1})) // Output: 15
}
