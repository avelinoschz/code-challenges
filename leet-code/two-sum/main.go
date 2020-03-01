package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("nums =", nums)
	fmt.Println("target =", 9)

	result := twoSum(nums, target)
	fmt.Println("output =", result)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("Comparing ", nums[i], nums[j])
			if (nums[i] + nums[j]) == target {
				fmt.Println("Found it!")
				return []int{i, j}
			}
		}
	}
	return []int{}
}
