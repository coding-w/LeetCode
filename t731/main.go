package main

func main() {

}

// MyCalendarTwo
// https://leetcode.cn/problems/my-calendar-ii/?envType=daily-question&envId=2025-01-03
// 每日一题，将重叠的部分记录下来
type MyCalendarTwo struct {
	booked   [][]int
	overlaps [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		booked:   [][]int{},
		overlaps: [][]int{},
	}
}

func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
	// 检查是否已有重叠部分
	for _, overlap := range this.overlaps {
		if overlap[0] < endTime && startTime < overlap[1] {
			return false
		}
	}
	for _, booked := range this.booked {
		if booked[0] < endTime && startTime < booked[1] {
			this.overlaps = append(this.overlaps, []int{max(booked[0], startTime), min(booked[1], endTime)})
		}
	}
	this.booked = append(this.booked, []int{startTime, endTime})
	return true
}
