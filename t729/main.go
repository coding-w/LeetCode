package main

func main() {
	c := Constructor()
	println(c.Book(37, 50))
	println(c.Book(4, 17))
	println(c.Book(35, 48))
}

// MyCalendar
// https://leetcode.cn/problems/my-calendar-i/?envType=daily-question&envId=2025-01-02
type MyCalendar struct {
	times [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		times: make([][]int, 0),
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	if len(this.times) == 0 {
		this.times = append(this.times, []int{startTime, endTime})
		return true
	}
	// 二分查找
	left, right := 0, len(this.times)-1
	mid := left + (right-left)/2
	for left <= right {
		if this.times[mid][0] <= startTime {
			if this.times[mid][1] > startTime {
				return false
			}
			left = mid + 1
		} else {
			if this.times[mid][0] < endTime {
				return false
			}
			right = mid - 1
		}
		mid = left + (right-left)/2
	}
	tmp := append([][]int{}, this.times[:mid]...)
	tmp = append(tmp, []int{startTime, endTime})
	this.times = append(tmp, this.times[mid:]...)
	return true
}
