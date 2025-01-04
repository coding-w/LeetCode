package main

import (
	"fmt"
	"sort"
)

func main() {
	c := Constructor()
	fmt.Println(c.Book(10, 20))
	fmt.Println(c.Book(50, 50))
	fmt.Println(c.Book(10, 40))
	fmt.Println(c.Book(5, 15))
	fmt.Println(c.Book(5, 10))
	fmt.Println(c.Book(25, 55))
}

// MyCalendarThree
// https://leetcode.cn/problems/my-calendar-iii/?envType=daily-question&envId=2025-01-04
// 每日一题
type MyCalendarThree struct {
	diffMap map[int]int
	maxK    int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		diffMap: make(map[int]int),
		maxK:    0,
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	// 差分数组的做法，记录区间的开始和结束
	this.diffMap[startTime]++
	this.diffMap[endTime]--

	// 提取并排序时间点
	var times []int
	for t := range this.diffMap {
		times = append(times, t)
	}
	sort.Ints(times) // 排序时间点

	cur := 0
	// 按照排序后的顺序遍历
	for _, t := range times {
		// 更新当前重叠数
		cur += this.diffMap[t]
		// 更新最大重叠数
		this.maxK = max(this.maxK, cur)
	}

	return this.maxK
}
