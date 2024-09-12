package main

import (
	"fmt"
)

// Deque represents a double-ended queue (deque) for integer indices.
type Deque struct {
	data []int
}

// PushBack adds an element to the back of the deque.
func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

// PopBack removes and returns the element from the back of the deque.
func (d *Deque) PopBack() {
	if len(d.data) > 0 {
		d.data = d.data[:len(d.data)-1]
	}
}

// Back returns the element at the back of the deque.
func (d *Deque) Back() int {
	return d.data[len(d.data)-1]
}

// PushFront adds an element to the front of the deque.
func (d *Deque) PushFront(val int) {
	d.data = append([]int{val}, d.data...)
}

// PopFront removes and returns the element from the front of the deque.
func (d *Deque) PopFront() {
	if len(d.data) > 0 {
		d.data = d.data[1:]
	}
}

// Front returns the element at the front of the deque.
func (d *Deque) Front() int {
	return d.data[0]
}

// Len returns the number of elements in the deque.
func (d *Deque) Len() int {
	return len(d.data)
}

// continuousSubarrays calculates the total number of continuous subarrays
// where the absolute difference between any two elements is <= 2.
func continuousSubarrays(nums []int) int64 {
	// Deques to track the maximum and minimum values
	maxDeque := &Deque{}
	minDeque := &Deque{}

	left := 0
	totalSubarrays := int64(0)

	for right := 0; right < len(nums); right++ {
		// Maintain the max deque (store indices in decreasing order of values)
		for maxDeque.Len() > 0 && nums[maxDeque.Back()] <= nums[right] {
			maxDeque.PopBack()
		}
		maxDeque.PushBack(right)

		// Maintain the min deque (store indices in increasing order of values)
		for minDeque.Len() > 0 && nums[minDeque.Back()] >= nums[right] {
			minDeque.PopBack()
		}
		minDeque.PushBack(right)

		// Move the left pointer to maintain the condition |nums[i1] - nums[i2]| <= 2
		for nums[maxDeque.Front()]-nums[minDeque.Front()] > 2 {
			left++
			// Remove elements that are out of the current window
			if maxDeque.Front() < left {
				maxDeque.PopFront()
			}
			if minDeque.Front() < left {
				minDeque.PopFront()
			}
		}

		// Number of valid subarrays ending at 'right' is (right - left + 1)
		totalSubarrays += int64(right - left + 1)
	}

	return totalSubarrays
}

func main() {
	fmt.Println(continuousSubarrays([]int{5, 4, 2, 4})) // Output: 8
	fmt.Println(continuousSubarrays([]int{1, 2, 3}))    // Output: 6
}
