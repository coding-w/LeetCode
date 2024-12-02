package main

import "fmt"

func main() {
	fmt.Println(getLargestOutlier([]int{6, -31, 50, -35, 41, 37, -42, 13}))
}

func getLargestOutlier(nums []int) int {
	sum := 0
	hmap := make(map[int]int)
	for i, num := range nums {
		sum += num
		hmap[num] = i
	}
	res := -1001
	for i := 0; i < len(nums); i++ {
		item := sum - nums[i]
		if item%2 == 0 {
			total := item / 2
			index, ok := hmap[total]
			if ok && i != index {
				res = max(res, nums[i])
			}
		}
	}
	return res
}
