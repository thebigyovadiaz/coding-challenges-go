package binarysearch

import "fmt"

// binarySearch performs binary search on a sorted array and returns the index of the target if found, or -1 otherwise.
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // Target found
		} else if arr[mid] < target {
			low = mid + 1 // Search the right half
		} else {
			high = mid - 1 // Search the left half
		}
	}

	return -1 // Target not found
}

func BinarySearch() {
	// Example usage:
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6

	result := binarySearch(sortedArray, target)

	if result != -1 {
		fmt.Printf("Binary-Serach >> Target %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Binary-Serach >> Target %d not found in the array\n", target)
	}
}
