package main

import (
	"fmt"
)

func twoSumEnhance(nums []int, target int) []int {
	obj := map[int]int{}
	for i := 0; i < len(nums); i++ {
		remainNumber := target - nums[i]
		fmt.Println(remainNumber, nums[i])
		if _, exists := obj[remainNumber]; exists {
			return []int{obj[remainNumber], i}
		} else {
			obj[nums[i]] = i
		}
	}

	return []int{}
}
func main() {
	result := twoSumEnhance([]int{1, 2, 4}, 6)
	fmt.Println(result)
}
