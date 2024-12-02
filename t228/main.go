package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))

}

// https://leetcode.cn/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	l, r := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			r++
			continue
		} else if r == l+1 {
			res = append(res, strconv.Itoa(nums[l]))
		} else {
			res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r-1]))
		}
		l = i
		r = i + 1
	}
	if r == l+1 {
		res = append(res, strconv.Itoa(nums[l]))
	} else {
		res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r-1]))
	}
	return res
}
