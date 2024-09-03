package main

import (
	"fmt"
	"sort"
)

func getMinimumCostMap(P, C, T []int) []map[int]int {
	// CM[i] is a map of costs to all destinations in descendants of i
	// CM[4][6] = 5 means, cost from 4th node to 6th node is 5
	// If CM[4][5] does not exists (key 5 is not prsent in the map stored in CM[4]),
	// then that does not mean there is no path from 4 to 5 (every node has path to every other node as it is a tree)
	// it just mean that either 5 is not marked as destination by any passenger or 5 is not a descendant of 4.
	CM := make([]map[int]int, len(T))
	for i := range CM {
		CM[i] = map[int]int{}
	}

	for _, n := range T {
		if n > 0 {
			// Calculate cost to this destination (n) from all it's parent (upto self)
			curN := n    // Self
			curCost := 0 // Cost to self is zero
			for curN > 0 {
				CM[curN][n] = curCost
				curCost += C[curN]
				curN = P[curN]
			}
		}
	}

	return CM
}

func minCostFromSourceToDest(s, d int, P, C []int, CM []map[int]int) int {
	cost := 0

	curS := s
	_, ok := CM[curS][d]
	for !ok {
		cost += C[curS] // Calculating cost by GOING UP to the common parent node of s and d
		curS = P[curS]
		if curS == 0 {
			// loop should never reach a point where curS becomes 0.
			// As each destination should be added in the cost map of node 1 i.e. CM[1][d] will always be true
			panic("destination set or cost map is not properly calculated.")
		}

		_, ok = CM[curS][d]
	}

	return cost + CM[curS][d] // cost of GOING UP to the common parent node of s and d + cost of GOING DOWN to the d from the common parent of s and d
}

// maxPassengers finds the maximum number of passengers that can be transported within the maximum travel distance K.
func maxPassengers(costs []int, K int) int {
	sort.Ints(costs)

	count := 0
	for _, cost := range costs {
		if cost <= K {
			K -= cost
			count++
		} else {
			break
		}
	}

	return count
}

func maxPeopleServed(N, K int, P, C, T []int) int {
	// Get the cost map for each node
	CM := getMinimumCostMap(P, C, T)

	// Get all the costs
	costs := []int{}
	for s, d := range T {
		if d > 0 {
			minCost := minCostFromSourceToDest(s, d, P, C, CM)
			costs = append(costs, CM[1][s]+minCost+CM[1][d]) // Cost of root (1) to source (s) + cost of source (s) to destination (d) + cost of destination (d) to root (1)

			// fmt.Printf("Minimum distance from %d to %d is %d\n", s, d, minCost) // Debug Statement
		}
	}

	// fmt.Println("Costs: ", costs) // Debug Statement
	return maxPassengers(costs, K)
}

func main() {
	// N := 2
	// K := 1
	// P := []int{0, 0, 1}
	// C := []int{0, 0, 2}
	// T := []int{0, 2, 1}
	// // Result: 0

	// N := 2
	// K := 2
	// P := []int{0, 0, 1}
	// C := []int{0, 0, 1}
	// T := []int{0, 2, 1}
	// // Result: 1

	// N := 5
	// K := 10
	// P := []int{0, 0, 1, 1, 2, 3}
	// C := []int{0, 0, 3, 2, 4, 1}
	// T := []int{0, 0, 0, 3, 0, 5}
	// // Result: 2

	N := 9
	K := 94
	P := []int{0, 0, 1, 1, 3, 3, 4, 4, 6, 6}
	C := []int{0, 0, 10, 15, 20, 10, 5, 25, 8, 12}
	T := []int{0, 0, 5, 6, 0, 8, 0, 9, 0, 4}
	// Result: 1

	result := maxPeopleServed(N, K, P, C, T)
	fmt.Println(result)
}
