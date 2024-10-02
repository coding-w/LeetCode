package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))
}

// https://leetcode.cn/problems/divide-two-integers/description/
func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	return dividend / divisor

}
