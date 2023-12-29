package mergeintervals

import (
	"fmt"
	"sort"
)

// Interval represents a time interval with a start and end.
type Interval struct {
	Start int
	End   int
}

// mergeIntervals merges overlapping intervals.
func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals based on their start times.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		previous := merged[len(merged)-1]

		if current.Start <= previous.End {
			// Merge overlapping intervals
			merged[len(merged)-1] = Interval{Start: previous.Start, End: maxI(current.End, previous.End)}
		} else {
			// Add non-overlapping interval
			merged = append(merged, current)
		}
	}

	return merged
}

// max returns the maximum of two integers.
func maxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MergeIntervals() {
	// Example usage:
	intervals := []Interval{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	mergedIntervals := mergeIntervals(intervals)

	fmt.Println("Merged Intervals:")
	for _, interval := range mergedIntervals {
		fmt.Printf("[%d, %d] ", interval.Start, interval.End)
	}
	// Output: Merged Intervals: [1, 6] [8, 10] [15, 18]
}
