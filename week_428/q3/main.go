package main

import "fmt"

func beautifulSplits(nums []int) int {
	n := len(nums)

	count := 0

	// 遍历所有可能的分割
	for i := 1; i < n-1; i++ {
		// 尝试将 nums 分割为 3 段
		nums1 := nums[:i]
		flag := true
		for j := i; j < n-1; j++ {
			nums2 := nums[i : j+1]
			nums3 := nums[j+1:]
			// 判断是否符合前缀条件
			if flag && isPrefix(nums1, nums2) {
				count += n - 1 - i
				break
			} else if isPrefix(nums2, nums3) {
				flag = false
				count++
			}
		}

	}

	return count
}

// 辅助函数，判断 nums1 是否是 nums2 的前缀
func isPrefix(nums1, nums2 []int) bool {
	if len(nums1) > len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 1, 2}
	result := beautifulSplits(nums)
	fmt.Println("Beautiful split count:", result)
}
