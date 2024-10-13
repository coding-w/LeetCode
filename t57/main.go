package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(insert([][]int{{1, 3}, {8, 10}}, []int{2, 6}))

}

// https://leetcode.cn/problems/insert-interval/description/
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}

	}
	return res
}
