package main

func main() {

}

// https://leetcode.cn/problems/majority-element/
// 计算数组中保留最多的元素
func majorityElement(nums []int) int {
	res := make([]int, 0)
	res = append(res, nums[0])
	for i := 1; i < len(nums); i++ {
		if len(res) > 0 && res[len(res)-1] != nums[i] {
			res = res[:len(res)-1]
		} else {
			res = append(res, nums[i])
		}
	}
	return res[0]

}
