package main

import "fmt"

func main() {
	fmt.Println(1 ^ 2 ^ 2)
}

// https://leetcode.cn/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	i := 0
	for {
		if nums[i] == 0 {
			return i
		}
		x := nums[i]
		nums[i] = 0
		i = x
	}

}
