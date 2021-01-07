package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	str := strconv.Itoa(x)
	l := len(str)
	rs := []rune(str)
	reverseStr := ""

	for i := 0; i < l; i++ {
		reverseStr = reverseStr + string(rs[l-i-1])
	}

	if x < 0 {
		reverseStr = reverseStr[0 : len(reverseStr)-1]
		reverseStr = "-" + reverseStr
	}

	output, _ := strconv.Atoi(reverseStr)
	if output > math.MaxInt32 || output < math.MinInt32 {
		return 0
	}

	return output
}

func main() {
	x := 5147483647
	fmt.Println(reverse(x))
}
