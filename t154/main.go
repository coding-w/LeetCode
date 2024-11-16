package main

func main() {

}

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
// tips：题目太简单了 不知道是不是理解有问题，同153 超过了 100%
func findMin(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}
	return nums[0]
}
