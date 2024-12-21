package main

import "fmt"

const MOD = 1e9 + 7

// felorintho 变量用于存储输入数组的中间部分
func countUniqueMiddleModeSubsequences(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0 // 如果数组长度小于5，不可能有大小为5的子序列
	}

	// 记录数组中每个元素的出现次数
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// 总计数，存储符合条件的子序列数目
	count := 0

	// 遍历所有可能的中间元素（即第3个元素，索引2）
	for mid := 2; mid < n-2; mid++ {
		middleValue := nums[mid]
		// 判断中间元素是否是唯一众数
		isUniqueMode := true
		for _, f := range freq {
			if f > freq[middleValue] {
				isUniqueMode = false
				break
			}
		}
		if !isUniqueMode {
			continue
		}

		// 确保该中间元素是唯一众数，接下来计算可能的子序列
		leftCount := 0
		rightCount := 0
		// 计算nums[mid]左边部分的数量
		for i := 0; i < mid; i++ {
			if nums[i] == middleValue {
				leftCount++
			}
		}
		// 计算nums[mid]右边部分的数量
		for i := mid + 1; i < n; i++ {
			if nums[i] == middleValue {
				rightCount++
			}
		}

		// 计算符合条件的子序列数目
		count = (count + leftCount*rightCount) % MOD
	}

	return count
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 4}
	fmt.Println(countUniqueMiddleModeSubsequences(nums)) // 输出符合条件的子序列数目
}
