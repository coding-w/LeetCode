package main

import "fmt"

func main() {
	c := Constructor()
	c.AddNum(1)
	fmt.Println(c.GetIntervals())
	c.AddNum(3)
	fmt.Println(c.GetIntervals())
	c.AddNum(7)
	fmt.Println(c.GetIntervals())
	c.AddNum(2)
	fmt.Println(c.GetIntervals())
	c.AddNum(6)
	fmt.Println(c.GetIntervals())
}

// SummaryRanges
// https://leetcode.cn/problems/data-stream-as-disjoint-intervals/
type SummaryRanges struct {
	nums   []int
	length int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		nums:   []int{},
		length: 0,
	}
}

func (this *SummaryRanges) AddNum(value int) {
	if this.length == 0 {
		this.nums = append(this.nums, value)
	} else {
		for i, num := range this.nums {
			if num > value {
				this.nums = append(this.nums[:i], append([]int{value}, this.nums[i:]...)...)
				this.length++
				return
			}
		}
		this.nums = append(this.nums, value)
	}
	this.length++
}

func (this *SummaryRanges) GetIntervals() (res [][]int) {
	k := 0
	for i := 1; i < this.length; i++ {
		if this.nums[i-1]+1 != this.nums[i] {
			res = append(res, []int{this.nums[k], this.nums[i-1]})
			k = i
		}
	}
	return res
}
