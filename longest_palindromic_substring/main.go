package main

import "fmt"

func checkPalindrome(s string, i, j int, max string) string {
	var sub string

	count := len(s)
	maxIndex := count - 1
	for i >= 0 && j <= maxIndex && s[i] == s[j] {
		sub = s[i : j+1]

		// 回文では中央の両脇が同じ文字になるので開始と終了のindexを一つずつずらす
		i--
		j++
	}

	if len(sub) > len(max) {
		return sub
	}

	return max
}

func longestPalindrome(s string) string {
	count := len(s)

	if count < 2 {
		return s
	}

	max := string(s[0])

	for i := 0; i < count; i++ {
		// xyxみたいな中央と両脇が異なる場合
		max = checkPalindrome(s, i, i, max)
		// abbcみたいな同じ文字が2文字並ぶ場合
		max = checkPalindrome(s, i, i+1, max)
	}

	return max
}

func main() {
	s := "ac"

	output := longestPalindrome(s)
	fmt.Println(output)
}
