package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)

	var medianIndex int
	if totalLen%2 == 0 {
		medianIndex = totalLen/2 + 1
	} else {
		medianIndex = (totalLen + 1) / 2
	}

	i, j, k := 0, 0, 0
	var prevMedian, median int
	for k < medianIndex {
		if j == len(nums2) || (i < len(nums1) && nums1[i] < nums2[j]) {
			prevMedian = median
			median = nums1[i]
			i++
		} else {
			prevMedian = median
			median = nums2[j]
			j++
		}
		k++
	}

	// fmt.Printf("i: %d, j: %d, median: %d\n", i, j, median)

	if totalLen%2 == 0 {
		return float64(prevMedian+median) / 2
	}

	return float64(median)
}

func main() {
	var nums1, nums2 []int

	nums1 = []int{1, 3}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Expected: 2.0

	nums1 = []int{1, 3}
	nums2 = []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Expected: 2.5

	nums1 = []int{1}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Expected: 1.5

	nums1 = []int{1}
	nums2 = []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Expected: 1

	nums1 = []int{}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Expected: 2
}
