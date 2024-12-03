package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countDigitOne(824883294))
}

// https://leetcode.cn/problems/number-of-digit-one/
// todo
func countDigitOne(n int) int {
	sb := strings.Builder{}
	for i := 1; i <= n; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	fmt.Println(sb.String())
	return strings.Count(sb.String(), "1")
}
