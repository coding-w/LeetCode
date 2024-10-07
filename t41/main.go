package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}

// https://leetcode.cn/problems/first-missing-positive/description/
func firstMissingPositive(nums []int) int {
	num := 0
	temp := 0
	kv := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			kv[nums[i]] = i
			if num < nums[i] {
				num = nums[i]
			}
		}
	}
	if num == 0 {
		return 1
	}
	for i := 1; i < num; i++ {
		if _, ok := kv[i]; !ok {
			temp = i
			break
		}
	}
	if temp == 0 {
		return num + 1
	}
	return temp

}
