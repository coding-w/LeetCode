package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

// https://leetcode.cn/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[v]; ok && i-j <= k {
			return true
		} else {
			m[v] = i
		}

	}
	return false

}
