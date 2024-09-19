package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	resMap := map[int]int{}
	for index, value := range nums {
		temp := target - value
		i, ok := resMap[temp]
		if ok {
			return []int{i, index}
		} else {
			resMap[value] = index
		}
	}
	return []int{}
}
