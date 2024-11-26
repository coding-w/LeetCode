package main

func main() {

}

// https://leetcode.cn/problems/minimum-size-subarray-sum/
// tips：滑动窗口
func minSubArrayLen(target int, nums []int) int {
	fast, slow, min_len, sum := 0, 0, len(nums)+1, 0
	for fast < len(nums) {
		sum += nums[fast]
		for sum >= target && fast >= slow {
			min_len = min(min_len, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
		fast++
	}
	if min_len == slow+1 {
		return 0
	}
	return min_len
}
