package main

import "fmt"

func main() {
	fmt.Println(countBits(11))
}

// https://leetcode.cn/problems/counting-bits/
// 找规律
// i		1
// 0		0
// 1		1
// 2		1
// 3		2
// 4		1
// 5		2
// 6		2
// 7		3
func countBits(n int) []int {
	nums := []int{0, 1, 1}
	if n < 3 {
		return nums[:n+1]
	}
	nums = make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	nums[2] = 1
	i, j := 3, 1
	m := 2
	for i <= n {
		if i == m*2 {
			nums[i] = 1
			m *= 2
			j = 1
			i++
			continue
		}
		nums[i] = nums[m] + nums[j]
		i++
		j++
	}
	return nums
}
