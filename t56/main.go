package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{2, 6}, {1, 3}, {8, 10}}))

}

// https://leetcode.cn/problems/merge-intervals/description/
// tips: 现将组合 根据intervals[x][0]进行排序 ==> intervals[x][0] < intervals[x+1][0]
// 只需要判断 intervals[x+1][1] 是否在 intervals[x] 区间内 即可
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 0; i < len(intervals); i++ {
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
