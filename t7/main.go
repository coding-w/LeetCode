package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))

}

// https://leetcode.cn/problems/reverse-integer/description/
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var res int64 = 0
	for x != 0 {
		res = res*10 + int64(x%10)
		x = x / 10

	}
	if res < int64(math.MaxInt32) && res > int64(math.MinInt32) {
		return int(res)
	}
	return 0
}
