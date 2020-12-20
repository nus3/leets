package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	firstIndex := 0
	secondIndex := 0

	for i := 0; i < numsLen; i++ {

		for j := i + 1; j < numsLen; j++ {
			if nums[i]+nums[j] == target {
				firstIndex = i
				secondIndex = j
				break
			}
		}

		if firstIndex != 0 && secondIndex != 0 {
			break
		}
	}

	return []int{firstIndex, secondIndex}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	output := twoSum(nums, target)
	fmt.Println(output)
}
