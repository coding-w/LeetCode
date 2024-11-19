package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(2147483647))
}

// https://leetcode.cn/problems/excel-sheet-column-title/
// tips: 类似二进制
func convertToTitle(columnNumber int) string {
	s := ""
	for columnNumber != 0 {
		s = string(rune('A'+(columnNumber-1)%26)) + s
		columnNumber = (columnNumber - 1) / 26
	}
	return s
}
