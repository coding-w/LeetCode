package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{1, 2, 3, 1}, 1))

}

// https://leetcode.cn/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
