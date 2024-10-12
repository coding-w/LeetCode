package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

// https://leetcode.cn/problems/jump-game/description/
func canJump(nums []int) bool {
	maxJ := nums[0]
	for i := 0; i < len(nums); i++ {
		if i > maxJ {
			return false
		}
		maxJ = max(maxJ, i+nums[i])
	}
	return true
}
