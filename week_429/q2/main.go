package main

import (
	"fmt"
	"sort"
)

// 超时....
func maxDistinctElements(nums []int, k int) int {
	// 对数组进行排序
	sort.Ints(nums)

	// 用 map 来存储已出现的元素
	seen := make(map[int]bool)

	// 遍历每个元素
	for _, num := range nums {
		// 先尝试不改变当前元素
		if !seen[num-k] {
			seen[num-k] = true
		} else {
			// 如果当前元素已存在，则尝试调整元素
			for delta := -k; delta <= k; delta++ {
				newNum := num + delta
				if !seen[newNum] {
					seen[newNum] = true
					break
				}
			}
		}
	}

	// 返回不同元素的数量
	return len(seen)
}

func main() {
	// 示例
	nums := []int{7, 8, 10, 10, 7, 6, 7}
	k := 1
	result := maxDistinctElements(nums, k)
	fmt.Println(result) // 输出：5
}
