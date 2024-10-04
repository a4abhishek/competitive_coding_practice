package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2

	imin, imax := 0, m

	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			// Increase i
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// Decrease i
			imax = i - 1
		} else {
			// i is perfect
			var maxOfLeft int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}

			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			var minOfRight int
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}

	return 0.0
}

func main() {
	var nums1, nums2 []int

	nums1 = []int{1, 3}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Expected: 2

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
