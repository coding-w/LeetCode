package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

// https://leetcode.cn/problems/top-k-frequent-elements/description/
// 时间太长....
func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	for _, n := range nums {
		hash[n]++
	}
	tmp := make([][]int, 100001)
	for item := range tmp {
		tmp[item] = make([]int, 0)
	}
	for key, v := range hash {
		tmp[v] = append(tmp[v], key)
	}
	res := make([]int, 0)
	for i := 100000; i > 0; i-- {
		if len(tmp[i]) > 0 {
			res = append(res, tmp[i]...)
			if len(res) >= k {
				break
			}
		}
	}
	return res[:k]
}
