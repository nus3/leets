package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)
	sort.Ints(arr)
	count := len(arr)

	var isOdd bool
	if count%2 == 0 {
		isOdd = false
	} else {
		isOdd = true
	}

	if count == 1 {
		return float64(arr[0])
	}

	index := count / 2
	if isOdd {
		return float64(arr[index])
	}

	return float64(arr[index-1]+arr[index]) / 2

}

func main() {
	// nums1 := []int{1, 3}
	nums1 := []int{1, 3}
	// nums2 := []int{2, 7}
	nums2 := []int{2}

	output := findMedianSortedArrays(nums1, nums2)
	fmt.Println(output)
}
