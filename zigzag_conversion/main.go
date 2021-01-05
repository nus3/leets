package main

import "fmt"

func convert(s string, numRows int) string {
	m := make(map[int]string)
	mIndex := 0
	isAdd := true

	if numRows == 1 {
		return s
	}

	for _, c := range s {
		str := string(c)
		m[mIndex] = m[mIndex] + str

		if mIndex == (numRows - 1) {
			isAdd = false
		}

		if mIndex == 0 {
			isAdd = true
		}

		if isAdd {
			mIndex++
		} else {
			mIndex--
		}
	}

	var output string

	for i := 0; i < numRows; i++ {
		output = output + m[i]
	}
	return output
}

func main() {
	s := "ABC"
	numRows := 2

	output := convert(s, numRows)
	fmt.Println(output)
}
