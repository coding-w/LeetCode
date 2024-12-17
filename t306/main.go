package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isAdditiveNumber("101"))
}

// https://leetcode.cn/problems/additive-number/
func isAdditiveNumber(num string) bool {

	n := len(num)
	mid := make([]int, 0)
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == n && len(mid) >= 3 {
			return true
		}
		for i := index; i < n; i++ {
			a := num[index : i+1]
			if a[0] == '0' && len(a) != 1 {
				break
			}
			b, _ := strconv.Atoi(a)
			if len(mid) < 2 || b == mid[len(mid)-1]+mid[len(mid)-2] {
				mid = append(mid, b)
				if dfs(i + 1) {
					return true
				}
				mid = mid[:len(mid)-1]
			}

		}

		return false
	}

	return dfs(0)
}
