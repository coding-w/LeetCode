package main

func main() {

}

// https://leetcode.cn/problems/minimum-operations-to-exceed-threshold-value-i/?envType=daily-question&envId=2025-01-14
func minOperations(nums []int, k int) int {
	res := 0
	for _, v := range nums {
		if v < k {
			res++
		}
	}
	return res
}
