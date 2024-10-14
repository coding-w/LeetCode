package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getPermutation(3, 4))
}

// https://leetcode.cn/problems/permutation-sequence/description/
// tips 康拓展开
func getPermutation(n int, k int) string {
	factorials := make([]int, n)
	nums := make([]int, n)
	sb := strings.Builder{}
	factorials[0] = 1
	for i := 1; i < n; i++ {
		factorials[i] = factorials[i-1] * i
	}

	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	k--
	for i := n - 1; i >= 0; i-- {
		index := k / factorials[i]
		k %= factorials[i]

		sb.WriteString(strconv.Itoa(nums[index]))
		nums = append(nums[:index], nums[index+1:]...)
	}

	return sb.String()
}
