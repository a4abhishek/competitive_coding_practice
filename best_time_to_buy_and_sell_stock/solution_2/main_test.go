package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, expected: 5},
		{prices: []int{7, 6, 4, 3, 1}, expected: 0},
		{prices: []int{1, 2, 3, 4, 5}, expected: 4},
		{prices: []int{2, 1, 2, 0, 1}, expected: 1},
		{prices: []int{3, 2, 6, 5, 0, 3}, expected: 4},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := maxProfit(test.prices)
			if result != test.expected {
				t.Errorf("maxProfit(%v) = %d; want %d", test.prices, result, test.expected)
			}
		})
	}
}
