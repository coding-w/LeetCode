package main

func main() {

}

// https://leetcode.cn/problems/split-the-array/?envType=daily-question&envId=2024-12-28
// 每日一题
func isPossibleToSplit(nums []int) bool {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
		if hash[v] > 2 {
			return false
		}
	}
	return true

}
