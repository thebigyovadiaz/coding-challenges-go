package knapsack

import "fmt"

// Knapsack represents an item with a weight and a value.
type Knapsack struct {
	Weight int
	Value  int
}

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// KnapsackDP solves the 0/1 Knapsack problem using dynamic programming.
func KnapsackDP(items []Knapsack, capacity int) int {
	n := len(items)

	// Create a 2D slice to store solutions to subproblems
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Build the table bottom-up
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if items[i-1].Weight <= w {
				// Include the current item in the knapsack
				include := items[i-1].Value + dp[i-1][w-items[i-1].Weight]
				// Exclude the current item from the knapsack
				exclude := dp[i-1][w]
				dp[i][w] = Max(include, exclude)
			} else {
				// The current item is too heavy to include
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func ResultKnapsack() {
	// Example usage:
	items := []Knapsack{
		{Weight: 2, Value: 3},
		{Weight: 1, Value: 2},
		{Weight: 3, Value: 5},
		{Weight: 4, Value: 4},
	}

	capacity := 5
	maxValue := KnapsackDP(items, capacity)

	fmt.Printf("KNAPSACK >> Maximum value achievable with a knapsack of capacity %d: %d\n", capacity, maxValue)
}
