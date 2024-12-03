package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{1, 2, 1, 2, 3}))
}

// https://leetcode.cn/problems/majority-element-ii/
// tips: 摩尔投票
func majorityElement(nums []int) []int {
	res := make([]int, 0)
	x1, count1, x2, count2 := 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if (count1 == 0 || nums[i] == x1) && x2 != nums[i] {
			count1++
			x1 = nums[i]
		} else if count2 == 0 || nums[i] == x2 {
			count2++
			x2 = nums[i]
		} else {
			count1--
			count2--
		}
	}
	// 校验是否正确
	count1, count2 = 0, 0
	for _, v := range nums {

		if v == x1 {
			count1++
		}
		if v == x2 {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		res = append(res, x1)
	}
	if count2 > len(nums)/3 && x1 != x2 {
		res = append(res, x2)
	}
	return res
}
