package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4, 2, 3, 3, 5, 7}))
}

func minimumOperations(nums []int) int {
	hmap := make(map[int]int)
	for _, n := range nums {
		hmap[n]++
	}
	for k, v := range hmap {
		if v == 1 {
			delete(hmap, k)
		}
	}
	res := 0
	for len(hmap) > 0 {
		res++
		if len(nums) < 3 {
			break
		}
		front := nums[:3]
		nums = nums[3:]
		for _, v := range front {
			if _, ok := hmap[v]; ok {
				hmap[v]--
				if hmap[v] == 1 {
					delete(hmap, v)
				}
			}
		}
	}

	//fmt.Println(nums)

	return res
}
