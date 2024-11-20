package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("ZY"))
}

// https://leetcode.cn/problems/excel-sheet-column-number/
// tips: 类似二进制数...
func titleToNumber(columnTitle string) int {
	i := 0
	ll := len(columnTitle)
	res := 0
	for i < len(columnTitle) {
		res += int(math.Pow(26, float64(ll-i-1))) * int(columnTitle[i]-'A'+1)
		i++
	}
	return res
}
