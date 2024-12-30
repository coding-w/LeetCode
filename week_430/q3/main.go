package main

import (
	"fmt"
)

func numberOfSubsequences(nums []int) int64 {
	n := len(nums)
	var count int64 = 0
	// p q r s
	// nums[p]*nums[r] == nums[q]*nums[s]		====> 	nums[p]/nums[q] == nums[s]/nums[r]
	hashMap := make(map[float64]int64)
	for r := 4; r < n; r++ {
		for p := r - 4; p >= 0; p-- {
			product := float64(nums[p]) / float64(nums[r-2])
			hashMap[product]++
		}
		for s := r + 2; s < n; s++ {
			product := float64(nums[s]) / float64(nums[r])
			count += hashMap[product]
		}
	}
	return count
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
	nums := []int{2, 19, 6, 13, 18, 13, 6, 6}
	result := numberOfSubsequences(nums)
	fmt.Println(result) // 输出 1
}
