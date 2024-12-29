package main

import (
	"fmt"
)

func numberOfSubsequences(nums []int) int64 {
	n := len(nums)
	count := 0

	// p q r s
	// 遍历所有符合条件的 q
	for q := 1; q < n-3; q++ {
		hashMap := make(map[int]int) // 用于存储 nums[p] * nums[q] 的值及其出现的次数

		// 枚举 p (在 q 前面，且满足 q - p > 1)
		for p := 0; p < q-1; p++ {
			product := nums[p] * nums[q]
			hashMap[product]++
		}

		if len(hashMap) == 0 {
			continue
		}

		// 枚举 r 和 s (在 q 后面，且满足 r - q > 1 和 s - r > 1)
		for r := q + 2; r < n-1; r++ {
			for s := r + 2; s < n; s++ {
				product := nums[r] * nums[s]
				// 判断是否在哈希表中
				if freq, exists := hashMap[product]; exists {
					count += freq // 匹配的组合数量
				}
			}
		}
	}

	return int64(count)
}

// 超时 ...
func numberOfSubsequences1(nums []int) int64 {
	n := len(nums)
	uniqueSequences := make(map[[4]int]struct{})
	for p := 0; p < n; p++ {
		for q := p + 2; q < n; q++ {
			for r := q + 2; r < n; r++ {
				for s := r + 2; s < n; s++ {
					if nums[p]*nums[r] == nums[q]*nums[s] {
						sequence := [4]int{p, q, r, s}
						uniqueSequences[sequence] = struct{}{}
					}
				}
			}
		}
	}

	return int64(len(uniqueSequences))
}

func main() {
	// 示例输入
	nums := []int{8, 2, 19, 6, 13, 18, 13, 6, 6}
	result := numberOfSubsequences(nums)
	fmt.Println(result) // 输出 1
}
