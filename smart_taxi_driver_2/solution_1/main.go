package main

import (
	"fmt"
)

const MAX_LOG = 17 // Tree height <= 2^17 (131072) i.e. just over maximum possible N as per constraints

var (
	N, K       int
	P, C, T, S []int

	depth  []int
	up     [][]int
	costUp [][]int
	dp     [][]int
)

func dfs(v int, parent int, d int, cost int) {
	depth[v] = d
	up[v][0] = parent
	costUp[v][0] = cost

	for i := 1; i < MAX_LOG; i++ {
		up[v][i] = up[up[v][i-1]][i-1]
		costUp[v][i] = costUp[v][i-1] + costUp[up[v][i-1]][i-1]
	}

	for i := 1; i <= N; i++ {
		if P[i] == v {
			dfs(i, v, d+1, C[i])
		}
	}
}

func preprocess(root int) {
	depth = make([]int, N+1)
	up = make([][]int, N+1)
	costUp = make([][]int, N+1)
	dp = make([][]int, N+1)
	for i := range up {
		up[i] = make([]int, MAX_LOG)
		costUp[i] = make([]int, MAX_LOG)
		dp[i] = make([]int, N+1)
		for j := range dp[i] {
			dp[i][j] = -1 // Initialize with -1 (indicating uncomputed)
		}
	}
	dfs(root, 0, 0, 0)
}

func lca(u, v int) int {
	if depth[u] < depth[v] {
		u, v = v, u
	}

	diff := depth[u] - depth[v]

	for i := 0; i < MAX_LOG; i++ {
		if (diff>>i)&1 != 0 {
			u = up[u][i]
		}
	}

	if u == v {
		return u
	}

	for i := MAX_LOG - 1; i >= 0; i-- {
		if up[u][i] != up[v][i] {
			u = up[u][i]
			v = up[v][i]
		}
	}

	return up[u][0]
}

func getCostToAncestor(u, ancestor int) int {
	totalCost := 0
	diff := depth[u] - depth[ancestor]

	for i := 0; i < MAX_LOG; i++ {
		if (diff>>i)&1 != 0 {
			totalCost += costUp[u][i]
			u = up[u][i]
		}
	}

	return totalCost
}

func minCostBetweenNodes(u, v int) int {
	// Check if already computed
	if dp[u][v] != -1 {
		return dp[u][v]
	}

	// Find LCA
	lcaNode := lca(u, v)

	// Calculate cost from u to v via LCA
	costU := getCostToAncestor(u, lcaNode)
	costV := getCostToAncestor(v, lcaNode)
	result := costU + costV

	// Store in DP table for future queries
	dp[u][v] = result
	dp[v][u] = result

	return result
}

func getSourceArrayFromDestinationArray(T []int) []int {
	// S is reverse of T
	// While T[i] = x gives the destination of the person at node i is x
	// S[i] = x gives the source (person standing at the node x) for which the destination is i
	S := make([]int, len(T))

	for s, d := range T[2:] { // S[0] = S[1] = 0
		if d > 0 { // do not touch S[0]
			S[d] = s + 2 // Because the slice starts from 2nd element
		}
	}

	return S
}

func maxPassengersFromChainWith(n int) int {
	// fmt.Printf("Calculating passengers from chain with %d\n", n) // Debug Statement

	// go to the end destination of the chain involving node n
	for T[n] > 0 {
		n = T[n]
	}

	// fmt.Printf("final destination of this chain is: %d\n", n) // Debug Statement

	// Backtrack to the first source of this chain
	maxPassengers := 0
	cost := 0
	curDest := n
	curSource := S[n]
	for curSource > 1 {
		// fmt.Printf("min cost between %d to %d is %d. Current cost is %d\n", curSource, curDest, minCostBetweenNodes(curSource, curDest), cost) // Debug Statement
		cost += minCostBetweenNodes(curSource, curDest)
		if cost <= K {
			// If cost is still manageable, then increment the passenger count
			maxPassengers++
		}

		// Mark passenger at this source as picked regardless of whether counted in maxPassengers
		// because this chain is already considered and we don't want to calculate this chain again from the caller of this function
		T[curSource] = 0 // Now there is no passenger waiting at node curSource

		// Backtrack
		curDest = curSource
		curSource = S[curSource]
	}

	return maxPassengers
}

func maxPeopleServed(N, K int, P, C, T []int) int {
	// Preprocess the tree for minimum distance calculation
	preprocess(1)

	maxPassengers := 0

	// Fix Self loop in transportation. i.e. passenger wants to go to i is already standing on i
	for s, d := range T {
		if s == d {
			maxPassengers = 1 // taxi driver could serve this passenger no matter what
		}

		T[s] = 0 // Removes self loop
	}

	// Get the source array (reverse of T)
	S = getSourceArrayFromDestinationArray(T)
	// fmt.Println("S: ", S) // Debug Statement

	// Find max passengers
	for s, d := range T {
		if d > 0 {
			passengersInCurrentChain := maxPassengersFromChainWith(s)
			// fmt.Printf("passengers from chain with %d is %d\n", s, passengersInCurrentChain) // Debug Statement
			if passengersInCurrentChain > maxPassengers {
				maxPassengers = passengersInCurrentChain
			}
		}
	}

	return maxPassengers
}

func main() {
	// N = 2
	// K = 1
	// P = []int{0, 0, 1}
	// C = []int{0, 0, 2}
	// T = []int{0, 2, 1}
	// // Result: 1

	// N = 2
	// K = 2
	// P = []int{0, 0, 1}
	// C = []int{0, 0, 1}
	// T = []int{0, 2, 1}
	// // Result: 1

	// N = 5
	// K = 10
	// P = []int{0, 0, 1, 1, 2, 3}
	// C = []int{0, 0, 3, 2, 4, 1}
	// T = []int{0, 0, 0, 3, 0, 5}
	// // Result: 1

	// N = 9
	// K = 94
	// P = []int{0, 0, 1, 1, 3, 3, 4, 4, 6, 6}
	// C = []int{0, 0, 10, 15, 20, 10, 5, 25, 8, 12}
	// T = []int{0, 0, 5, 6, 0, 8, 0, 9, 0, 4}
	// // Result: 2

	// N = 9
	// K = 90
	// P = []int{0, 0, 1, 1, 3, 3, 4, 4, 6, 6}
	// C = []int{0, 0, 10, 15, 20, 10, 5, 25, 8, 12}
	// T = []int{0, 0, 5, 4, 9, 6, 7, 0, 0, 0}
	// // Result: 2

	N = 9
	K = 100
	P = []int{0, 0, 1, 1, 3, 3, 4, 4, 6, 6}
	C = []int{0, 0, 10, 15, 20, 10, 5, 25, 8, 12}
	T = []int{0, 0, 5, 4, 9, 6, 7, 0, 0, 0}
	// Result: 3

	result := maxPeopleServed(N, K, P, C, T)
	fmt.Println(result)
}
