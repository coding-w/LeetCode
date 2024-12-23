package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
}

// https://leetcode.cn/problems/power-of-three/
// 1162261467 是在范围内最大的幂
// return n > 0 && 1162261467%n == 0
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1

}
