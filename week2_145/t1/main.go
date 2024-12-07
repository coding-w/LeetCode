package main

import "sort"

func main() {

}

func minOperations(nums []int, k int) int {

	sort.Ints(nums)
	if nums[0] < k {
		return -1
	}
	res := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			res++
			if nums[i] < k {
				break
			}
		}
	}
	if nums[0] != k {
		res++
	}
	return res
}
