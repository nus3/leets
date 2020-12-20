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

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		idx, ok := m[target-num]
		if ok {
			return []int{idx, i}
		}

		// NOTE: mapのindexに値を埋め込み、target-numした時に合致したindexがあればそん時のvalueとindexを返す
		m[num] = i
	}

	return nil
}

func main() {
	// nums := []int{3, 2, 4}
	nums := []int{11, 15, 2, 7}
	// target := 6
	target := 9
	output := twoSum2(nums, target)
	fmt.Println(output)
}
