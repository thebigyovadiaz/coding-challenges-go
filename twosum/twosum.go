package twosum

import "fmt"

func twoSum(nums []int, target int) []int {
	complementMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := complementMap[complement]; found {
			return []int{index, i}
		}

		complementMap[num] = i
	}

	return nil
}

func Twosum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("TWO-SUM >>", result)
}
