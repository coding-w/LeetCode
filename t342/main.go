package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfFour(18))
}

// https://leetcode.cn/problems/power-of-four/
func isPowerOfFour(n int) bool {
	if n < 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%4 != 0 {
		return false
	}
	return isPowerOfFour(n / 4)

}
