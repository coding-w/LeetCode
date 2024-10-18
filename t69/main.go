package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(2))
}

// https://leetcode.cn/problems/sqrtx/description/
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}
