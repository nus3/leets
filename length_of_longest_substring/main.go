package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	max := 0
	// charの開始位置をmapで格納
	m := make(map[rune]int)
	startIndex := 0

	for currentIndex, char := range s {
		if startIndex < m[char] {
			startIndex = m[char]
		}

		count := currentIndex - startIndex + 1
		if max < count {
			max = count
		}

		// 同じ文字列があった場合はその文字から次の文字よりスタートする
		m[char] = currentIndex + 1
	}

	return max
}

func main() {
	// s := "abcabcbbb"
	// s := "bbbbb"
	// s := "pwwkew"
	s := "dvdf"
	// s := ""
	output := lengthOfLongestSubstring(s)
	fmt.Println("result")
	fmt.Println(output)
}
