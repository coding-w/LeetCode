package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

// https://leetcode.cn/problems/climbing-stairs/description/
// tips: 第n个台阶= (n-1级+1) + (n-2 级 + 2)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}
